package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)
type StrArray []string

var supportedLanguages = StrArray{"python", "go"}

func main() {
	log.Fatal(http.ListenAndServe(":8000", getRouter()))
}

func getRouter () *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/test/{language}/{uuid}", RunTest).Methods("GET")
	router.HandleFunc("/", Welcome).Methods("GET")
	return router
}
