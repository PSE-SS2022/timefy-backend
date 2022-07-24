package handlers

import (
	"net/http"

	"github.com/PSE-SS2022/timefy-backend/cmd/auth"
)

func SignUpUser(response http.ResponseWriter, request *http.Request) {
	if auth.IsAllowed(request) {
		// Do something
		response.WriteHeader(http.StatusOK)
		return
	}
	response.WriteHeader(http.StatusUnauthorized)
}

func SendFCMToken(response http.ResponseWriter, request *http.Request) {
	if auth.IsAllowed(request) {
		// Do something
		response.WriteHeader(http.StatusOK)
		return
	}
	response.WriteHeader(http.StatusUnauthorized)
}
