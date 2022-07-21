package handlers

import (
	"context"
	"net/http"

	"github.com/PSE-SS2022/timefy-backend/internal/repos"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
)

func DeleteUser(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id := params["id"]
	repos.GetCollection("users").DeleteOne(context.TODO(), bson.D{{Key: "id", Value: id}})
}
