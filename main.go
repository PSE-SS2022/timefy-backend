package main

import (
	"log"
	"net/http"

	"github.com/PSE-SS2022/timefy-backend/cmd/auth"
	"github.com/PSE-SS2022/timefy-backend/cmd/handlers"
	"github.com/PSE-SS2022/timefy-backend/internal/repos"
	"github.com/PSE-SS2022/timefy-backend/internal/server"
	"github.com/gorilla/mux"
)

func main() {
	if err := repos.InitDB(); err != nil {
		panic("Error initializing the Database, error:" + err.Error())
	}
	auth.GetEnforcer() // init enforcer
	router := mux.NewRouter()

	fs := http.FileServer(http.Dir("web"))
	http.Handle("/web/", http.StripPrefix("/web/", fs))

	http.Handle("/", router)
	router.HandleFunc("/", handlers.HomePageHandler)
	router.HandleFunc("/login", handlers.LoginPageHandler).Methods("GET")
	router.HandleFunc("/login", auth.LoginHandler).Methods("POST")

	router.HandleFunc("/register", handlers.RegisterPageHandler).Methods("GET")
	router.HandleFunc("/register", handlers.RegisterAdmin).Methods("POST")

	log.Println("All handlers set and ready to listen")
	server.Start()
}
