package main

import (
	"log"
	"net/http"

	"github.com/PSE-SS2022/timefy-backend/cmd/auth"
	"github.com/PSE-SS2022/timefy-backend/cmd/handlers"
	"github.com/PSE-SS2022/timefy-backend/internal/server"

	"github.com/gorilla/mux"
)

// Functions for handling pagecalls like localhost:8080/login
func main() {
	/*if err := repos.InitDB(); err != nil {
		log.Fatal("Error initializing the Database, error:" + err.Error())
		return
	}*/
	enforcer := auth.SetUpRBAC()
	router := mux.NewRouter()
	router.Use(auth.Middleware(&auth.Authorize{Enforcer: enforcer}))
	fs := http.FileServer(http.Dir("web"))
	http.Handle("/web/", http.StripPrefix("/web/", fs))

	http.Handle("/", router)
	router.HandleFunc("/", handlers.HomePageHandler)
	router.HandleFunc("/login", handlers.LoginPageHandler)

	log.Println("All handlers set and ready to listen")
	server.Start()
}
