package handlers

import (
	"context"
	"fmt"
	"html/template"
	"net/http"

	"github.com/PSE-SS2022/timefy-backend/cmd/auth"
	"github.com/PSE-SS2022/timefy-backend/cmd/encryption"
	"github.com/PSE-SS2022/timefy-backend/internal/helpers"
	"github.com/PSE-SS2022/timefy-backend/internal/models"
	"github.com/PSE-SS2022/timefy-backend/internal/repos"
	"github.com/PSE-SS2022/timefy-backend/web"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func HomePageHandler(response http.ResponseWriter, request *http.Request) {
	if auth.IsAuthenticatedWithSession(request) {
		t, err := web.ParseFiles(helpers.JoinPaths("web", "templates", "index.gohtml"),
			helpers.JoinPaths("web", "templates", "base.tmpl"))
		if err != nil {
			fmt.Println(err)
		}
		t.Execute(response, GetReports())
	} else {
		response.Header().Set("Content-Type", "text/html")
		response.WriteHeader(http.StatusUnauthorized)
		response.Write([]byte(`<script>window.location.href = "/login";</script>`))
	}
}

func LoginPageHandler(response http.ResponseWriter, request *http.Request) {
	t, _ := web.ParseFiles(helpers.GetTemplate("login.gohtml"),
		helpers.GetTemplate("base.tmpl"))
	t.Execute(response, nil)
}

func RegisterPageHandler(response http.ResponseWriter, request *http.Request) {
	t, _ := template.ParseFiles(helpers.GetTemplate("register.gohtml"), helpers.GetTemplate("base.tmpl"))
	t.Execute(response, nil)
}

func RegisterHandler(response http.ResponseWriter, request *http.Request) {
	collection := repos.GetCollection("admins")
	request.ParseForm()

	// Get data the User typen into the fields
	firstName := request.FormValue("firstname")
	lastName := request.FormValue("lastname")
	email := request.FormValue("email")

	_firstName := !helpers.IsEmpty(firstName)
	_lastName := !helpers.IsEmpty(lastName)
	_email := !helpers.IsEmpty(email)
	if _firstName && _lastName && _email {
		result := collection.FindOne(context.TODO(), bson.D{{Key: "Email", Value: email}})
		if result.Err() != nil {
			usr := models.User{ID: primitive.NewObjectID(), FirstName: firstName, LastName: lastName, Email: email, Tags: []models.Tag{}, Roles: map[string]string{}, AmountWarnings: 0}
			collection.InsertOne(context.TODO(), usr)
			http.Redirect(response, request, "/?success", 302)
		} else {
			http.Redirect(response, request, "/register?wrong", 302)
			return
		}
	} else {
		http.Redirect(response, request, "/register?empty", 302)
		return
	}
}

func RegisterAdmin(response http.ResponseWriter, request *http.Request) {
	collection := repos.GetCollection("admins")
	request.ParseForm()
	email := request.FormValue("email")
	password := request.FormValue("password")
	firstName := encryption.Encrypt(request.FormValue("firstName"))
	lastName := encryption.Encrypt(request.FormValue("lastName"))
	var adminModel models.Admin
	// Look if the entered Mail is already used
	err := collection.FindOne(context.TODO(), bson.D{{Key: "Email", Value: email}}).Decode(&adminModel)
	// If not found (throws exception/error) then we can proceed, or if found but found one is not same admintype as found one we proceed
	if err != nil {
		// Generate the hashed password with 14 as salt
		hash, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
		newAdmin := models.Admin{ID: primitive.NewObjectID(), FirstName: firstName, LastName: lastName, Email: email, Password: string(hash), Role: models.ADMIN_ROLE}
		newAdmin.Role = models.ADMIN_ROLE
		collection.InsertOne(context.TODO(), newAdmin)
		http.Redirect(response, request, "/login", 302)
	} else {
		panic("admin exists already")
	}
}

func GetReports() []models.ExtendedReport {
	return demoReports
}

// TODO: need to add something in front of id as id may not start with an number --> jquery error
var demoReports = []models.ExtendedReport{
	models.ExtendedReport{"i" + primitive.NewObjectID().Hex(), "Abdullah#123", "Abdullah", "Yildirim", "21.05.2022", "1", "Mittagessen", "Hier treffen wir uns zum Mittagessen in der Mensa"},
	models.ExtendedReport{"i" + primitive.NewObjectID().Hex(), "Talip#124", "Talip", "Göksu", "19.05.2022", "2", "Fußball", "Hi, wer hat Lust auf Fußball"},
	models.ExtendedReport{"i" + primitive.NewObjectID().Hex(), "Barrack#125", "Barrack", "Obama", "10.05.2022", "3", "Murriicaa", "Murriicaaaaaa"},
}
