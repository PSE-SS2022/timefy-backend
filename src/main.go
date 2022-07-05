package main

import (
	"fmt"
	"log"
	"net/http"
	"timefy-backend/common"

	"github.com/gorilla/mux"
)

var router = mux.NewRouter()

// Functions for handling pagecalls like localhost:8080/login
func main() {
	fmt.Println("Starting up Server")
	/*if err := repos.InitDB(); err != nil {
		log.Fatal("Error initializing the Database, error:" + err.Error())
		return
	}*/
	fs := http.FileServer(http.Dir("website"))
	http.Handle("/website/", http.StripPrefix("/website/", fs))
	http.Handle("/", router)
	router.HandleFunc("/", common.HomePageHandler)
	router.HandleFunc("/login", common.LoginPageHandler)
	log.Println("All handlers set and ready to listen")
	http.ListenAndServe(":80", nil)
}
