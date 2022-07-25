package handlers

import (
	"context"
	"net/http"

	"github.com/PSE-SS2022/timefy-backend/cmd/auth"
	"github.com/PSE-SS2022/timefy-backend/internal/helpers"
	"github.com/PSE-SS2022/timefy-backend/internal/models"
	"github.com/PSE-SS2022/timefy-backend/internal/repos"
	"go.mongodb.org/mongo-driver/bson"
)

func SignUpUser(response http.ResponseWriter, request *http.Request) {
	if auth.IsAllowed(request, false) {
		collection := repos.GetCollection("admins")
		request.ParseForm()
		id := request.FormValue("id")
		firstName := request.FormValue("firstname")
		lastName := request.FormValue("lastname")
		email := request.FormValue("email")

		_id := !helpers.IsEmpty(id)
		_firstName := !helpers.IsEmpty(firstName)
		_lastName := !helpers.IsEmpty(lastName)
		_email := !helpers.IsEmpty(email)
		if _id && _firstName && _lastName && _email {
			result := collection.FindOne(context.TODO(), bson.D{{Key: "Email", Value: email}})
			if result.Err() != nil {
				usr := models.User{ID: id, FirstName: firstName, LastName: lastName, Email: email, Tags: []models.Tag{}, Roles: map[string]string{}, AmountWarnings: 0}
				collection.InsertOne(context.TODO(), usr)
				response.WriteHeader(http.StatusOK)
				return
			} else {
				response.WriteHeader(http.StatusBadRequest)
				return
			}
		} else {
			response.WriteHeader(http.StatusBadRequest)
			return
		}
	}
	response.WriteHeader(http.StatusUnauthorized)
}

func SendFCMToken(response http.ResponseWriter, request *http.Request) {
	if auth.IsAllowed(request, false) {
		// Do something
		response.WriteHeader(http.StatusOK)
		return
	}
	response.WriteHeader(http.StatusUnauthorized)
}
