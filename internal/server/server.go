package server

import (
	"log"
	"net/http"
)

func Start() {
	log.Fatal(http.ListenAndServe(":80", nil))
}
