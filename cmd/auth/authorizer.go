package auth

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/casbin/casbin"

	"github.com/PSE-SS2022/timefy-backend/internal/database"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

type Authorize struct {
	Enforcer *casbin.Enforcer
}

func (a *Authorize) HasPermission(userID, action, asset string) bool {
	user, ok := database.UserRepositoryInstance.GetUserById(userID)
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

func getTokenCookie(token string) *http.Cookie {
	return &http.Cookie{
		Name:  "token",
		Value: token,
		Path:  "/",
	}
}

func getTokenFromCookie(request *http.Request) string {
	cookie, err := request.Cookie("token")
	if err != nil {
		return ""
	}
	return cookie.Value
}

// To check if the user is authenticated
func IsAuthenticated(authClient *auth.Client, request *http.Request) (bool, error) {
	idToken := getTokenFromCookie(request)
	if idToken == "" {
		return false, errors.New("token not set or invalid")
	}

	if !authenticate(authClient, idToken) {
		return false, errors.New("invalid token")
	}
	return true, nil
}

func authenticate(authClient *auth.Client, idToken string) bool {
	_, err := authClient.VerifyIDToken(context.Background(), idToken)
	return err == nil
}

func SetupFirebase() *auth.Client {
	opt := option.WithCredentialsFile(os.Getenv("GOOGLE_APPLICATION_CREDENTIALS"))
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		panic(fmt.Errorf("error initializing app: %v", err))
	}

	//Firebase Auth
	auth, err := app.Auth(context.Background())
	if err != nil {
		panic(fmt.Errorf("firebase load error"))
	}

	return auth
}

func SetUpRBAC() *casbin.Enforcer {
	enforcer, err := casbin.NewEnforcerSafe("./rbac_model.conf", "./rbac_policy.csv")
	if err != nil {
		panic(fmt.Errorf("failed to create enforcer: %v", err))
	}
	return enforcer
}

// Function for setting the Cookie
func SetCookie(jwtToken string, response http.ResponseWriter) {
	http.SetCookie(response, getTokenCookie(jwtToken))
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

func LogIn(authClient *auth.Client, request *http.Request) (bool, error) {
	authorizationToken := request.Header.Get("Authorization")
	idToken := strings.TrimSpace(strings.Replace(authorizationToken, "Bearer", "", 1))
	if idToken == "" {
		return false, errors.New("no token provided")
	}

	if !authenticate(authClient, idToken) {
		return false, errors.New("invalid token")
	}
	return true, nil
}
