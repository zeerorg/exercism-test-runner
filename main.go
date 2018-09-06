package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/test/{language}/{uuid}", RunTest).Methods("GET")
	router.HandleFunc("/", Welcome).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}
