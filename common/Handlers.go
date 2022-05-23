package common

import (
	"html/template"
	"net/http"
	"timefy-backend/models"
)

type TemplateStruct struct {
	Reports []models.Report
}

func HomePageHandler(response http.ResponseWriter, request *http.Request) {
	t, _ := template.ParseFiles("website/index.html")
	t.Execute(response, nil)
}

func LoginPageHandler(response http.ResponseWriter, request *http.Request) {
	t, _ := template.ParseFiles("website/login.html")
	t.Execute(response, nil)
}
