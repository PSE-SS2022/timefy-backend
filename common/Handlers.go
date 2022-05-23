package common

import (
	"fmt"
	"html/template"
	"net/http"
	"timefy-backend/repos"
)

func HomePageHandler(response http.ResponseWriter, request *http.Request) {
	t, err := template.ParseFiles("website/index.gohtml", "website/base.tmpl", "website/footer.tmpl")
	if err != nil {
		fmt.Println(err)
	}
	t.Execute(response, repos.GetReports())
}

func LoginPageHandler(response http.ResponseWriter, request *http.Request) {
	t, _ := template.ParseFiles("website/login.html")
	t.Execute(response, nil)
}
