package main

import (
	"log"
	"net/http"
	"timefy-backend/common"

	"github.com/gorilla/mux"
)

var router = mux.NewRouter()

// Functions for handling pagecalls like localhost:8080/login
func main() {
	/*if err := repos.InitDB(); err != nil {
		log.Fatal("Error initializing the Database, error:" + err.Error())
		return
	}*/
	router.HandleFunc("/", common.HomePageHandler)
	http.Handle("/", router)
	log.Println("All handlers set and ready to listen")
	http.ListenAndServe(":80", nil)
}
