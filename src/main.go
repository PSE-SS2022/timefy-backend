package main

import (
	"log"
	"net/http"

	"github.com/PSE-SS2022/timefy-backend/src/handlers"

	"github.com/gorilla/mux"
)

var router = mux.NewRouter()

// Functions for handling pagecalls like localhost:8080/login
func main() {
	/*if err := repos.InitDB(); err != nil {
		log.Fatal("Error initializing the Database, error:" + err.Error())
		return
	}*/
	fs := http.FileServer(http.Dir("website"))
	http.Handle("/website/", http.StripPrefix("/website/", fs))

	http.Handle("/", router)
	router.HandleFunc("/", handlers.HomePageHandler)
	router.HandleFunc("/login", handlers.LoginPageHandler)

	log.Println("All handlers set and ready to listen")
	http.ListenAndServe(":80", nil)
}
