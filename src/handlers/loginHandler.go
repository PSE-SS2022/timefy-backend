package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"timefy-backend/src/repos"
)

func HomePageHandler(response http.ResponseWriter, request *http.Request) {
	t, err := template.ParseFiles("website/templates/index.gohtml", "website/templates/base.tmpl")
	if err != nil {
		fmt.Println(err)
	}
	t.Execute(response, repos.GetReports())
}

func LoginPageHandler(response http.ResponseWriter, request *http.Request) {
	t, _ := template.ParseFiles("website/templates/login.gohtml", "website/templates/base.tmpl")
	t.Execute(response, nil)
}
