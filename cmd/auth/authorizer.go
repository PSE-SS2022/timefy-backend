package auth

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/casbin/casbin"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"

	"github.com/PSE-SS2022/timefy-backend/internal/models"
	"github.com/PSE-SS2022/timefy-backend/internal/repos"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

type Authorize struct {
	Enforcer *casbin.Enforcer
}

type adminLoginCredentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

var authorizer Authorize
var authClient *auth.Client

var sessions = map[string]session{}

type session struct {
	email  string
	expiry time.Time
}

func LoginHandler(response http.ResponseWriter, request *http.Request) {
	if IsAuthenticatedWithSession(request) {
		response.Write([]byte(`<script>window.location.href = "/";</script>`))
		return
	} else {
		request.ParseForm()
		if len(request.PostForm) > 0 {
			if AuthenticateWithEmailAndPassword(request) {
				sessionToken := uuid.NewString()
				expiresAt := time.Now().Add(24 * time.Hour)

				// Set the token in the session map, along with the user whom it represents
				sessions[sessionToken] = session{
					email:  request.FormValue("email"),
					expiry: expiresAt,
				}
				SetCookie("session_token", sessionToken, expiresAt, response)
				response.Header().Set("Content-Type", "text/html")
				response.WriteHeader(http.StatusAccepted)
				response.Write([]byte(`<script>window.location.href = "/";</script>`))
			} else {
				response.Header().Set("Content-Type", "text/html")
				response.WriteHeader(http.StatusUnauthorized)
				response.Write([]byte(`<script>window.location.href = "/login";</script>`))
			}
		} else {
			response.Header().Set("Content-Type", "text/html")
			response.WriteHeader(http.StatusBadRequest)
			response.Write([]byte(`<script>window.location.href = "/login";</script>`))
		}
	}
}

func AuthenticateWithEmailAndPassword(request *http.Request) bool {
	email := request.FormValue("email")
	password := request.FormValue("password")
	if len(email) != 0 && len(password) != 0 {
		collection := repos.GetCollection("admins")
		var admin adminLoginCredentials
		err := collection.FindOne(context.TODO(), bson.D{{Key: "Email", Value: email}}).Decode(&admin)
		if err != nil {
			return false
		}
		err = bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(password))
		return err == nil
	} else {
		return false
	}
}

func (s session) isExpired() bool {
	return s.expiry.Before(time.Now())
}

func IsAuthenticatedWithSession(request *http.Request) bool {
	token := GetCookie("session_token", request)
	if token == "" {
		return false
	}
	session, ok := sessions[token]
	if !ok || session.isExpired() {
		delete(sessions, token)
		return false
	}
	return true
}

func GetMailFromSession(request *http.Request) string {
	token := GetCookie("session_token", request)
	if token == "" {
		return ""
	}
	session, ok := sessions[token]
	if !ok {
		return ""
	}
	println(session.email)
	return session.email
}

func Logout(response http.ResponseWriter, request *http.Request) {
	token := GetCookie("session_token", request)
	if token == "" {
		response.WriteHeader(http.StatusUnauthorized)
		return
	}
	delete(sessions, token)
	SetCookie("session_token", "", time.Now(), response)
	http.SetCookie(response, &http.Cookie{
		Name:    "session_token",
		Value:   "",
		Expires: time.Now(),
	})
}

func GetToken(request *http.Request) string {
	token := getTokenFromCookie(request)
	if token == "" {
		authorizationToken := request.Header.Get("Authorization")
		idToken := strings.TrimSpace(strings.Replace(authorizationToken, "Bearer", "", 1))
		return idToken
	}
	return token
}

func IsAllowed(request *http.Request, isAdmin bool) bool {
	authenticated := false
	authorized := false
	if isAdmin {
		authenticated = IsAuthenticatedWithSession(request)
		if !authenticated {
			return false
		}
		admin, exists := models.GetAdminByMail(GetMailFromSession(request))
		if !exists {
			return false
		}
		authorized = admin.Role == "admin" || admin.Role == "mod"
	} else {
		authenticated, _ = IsAuthenticatedWithBearer(request)
		if !authenticated {
			return false
		}
		authorized = authorizer.HasPermission(getUserIdByToken(getTokenFromCookie(request)), request.Method, request.URL.Path, false)
	}
	return authenticated && authorized
}

func (a *Authorize) HasPermission(id primitive.ObjectID, action, asset string, isAdmin bool) bool {
	if isAdmin {
		admin, ok := models.GetAdminById(id)
		if !ok {
			return false
		}
		return a.Enforcer.Enforce(admin.Role, asset, action)
	} else {
		user, ok := models.GetUserByID(id)
		if !ok {
			return false
		}
		for _, role := range user.Roles {
			if a.Enforcer.Enforce(role, asset, action) {
				return true
			}
		}
		return false
	}
}

func getUserIdByToken(token string) primitive.ObjectID {
	fireBaseUser, err := authClient.GetUser(context.Background(), token)
	if err != nil {
		return primitive.NilObjectID
	}
	user, ok := models.GetUserByMail(fireBaseUser.Email)
	if !ok {
		return primitive.NilObjectID
	}
	return user.ID
}

// To check if the user is authenticated
func IsAuthenticatedWithBearer(request *http.Request) (bool, error) {
	idToken := GetToken(request)
	if idToken == "" {
		return false, errors.New("token not set or invalid")
	}

	if !authenticate(idToken) {
		return false, errors.New("invalid token")
	}
	return true, nil
}

func authenticate(idToken string) bool {
	_, err := authClient.VerifyIDToken(context.TODO(), idToken)
	return err == nil
}

func SetupFirebase() {
	opt := option.WithCredentialsFile(os.Getenv("GOOGLE_APPLICATION_CREDENTIALS"))
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		panic(fmt.Errorf("error initializing app: %v", err))
	}

	//Firebase Auth
	auth, err := app.Auth(context.TODO())
	if err != nil {
		panic(fmt.Errorf("firebase load error"))
	}
	authClient = auth
}

func GetEnforcer() *casbin.Enforcer {
	if authorizer.Enforcer == nil {
		setUpRBAC()
	}
	return authorizer.Enforcer
}

func setUpRBAC() {
	enforcer, err := casbin.NewEnforcerSafe("configs/rbac_model.conf", "configs/rbac_policy.csv")
	if err != nil {
		panic(fmt.Errorf("failed to create enforcer: %v", err))
	}
	authorizer.Enforcer = enforcer
}

// Function for setting the Cookie
func SetCookie(name, value string, expires time.Time, response http.ResponseWriter) {
	http.SetCookie(response, getTokenCookie(name, value, expires))
}

func getTokenCookie(name, value string, expires time.Time) *http.Cookie {
	return &http.Cookie{
		Name:    name,
		Value:   value,
		Expires: expires,
	}
}

func GetCookie(name string, request *http.Request) string {
	cookie, err := request.Cookie(name)
	if err != nil {
		return ""
	}
	return cookie.Value
}

func getTokenFromCookie(request *http.Request) string {
	cookie, err := request.Cookie("token")
	if err != nil {
		return ""
	}
	return cookie.Value
}

// Function for deletion of the Cookie
func ClearCookie(response http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   "cookie",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(response, cookie)
}
