package handlers

import (
	"fmt"
	"net/http"

	"github.com/PSE-SS2022/timefy-backend/cmd/auth"
	"github.com/PSE-SS2022/timefy-backend/internal/helpers"
	"github.com/PSE-SS2022/timefy-backend/internal/models"
	"github.com/PSE-SS2022/timefy-backend/web"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func HomePageHandler(response http.ResponseWriter, request *http.Request) {
	if auth.IsAllowed(request, true) {
		t, err := web.ParseFiles(helpers.JoinPaths("web", "templates", "index.gohtml"),
			helpers.JoinPaths("web", "templates", "base.tmpl"))
		if err != nil {
			fmt.Println(err)
		}
		response.WriteHeader(http.StatusOK)
		t.Execute(response, GetReports())
	} else {
		response.Header().Set("Content-Type", "text/html")
		response.WriteHeader(http.StatusUnauthorized)
		response.Write([]byte(`<script>window.location.href = "/login";</script>`))
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
