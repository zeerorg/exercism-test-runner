package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func Welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Exercism test server is running")
}

func RunTest(w http.ResponseWriter, r *http.Request) {
	language := mux.Vars(r)["language"]
	uuid := mux.Vars(r)["uuid"]
	success := make(chan bool)
	go GetAsyncSolution(language, uuid, success)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(map[string]bool{
		"result": <-success,
	})
}
