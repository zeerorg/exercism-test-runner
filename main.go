package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os/exec"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/test/{language}/{uuid}", RunTest).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func RunTest(w http.ResponseWriter, r *http.Request) {
	language := mux.Vars(r)["language"]
	uuid := mux.Vars(r)["uuid"]
	success := make(chan bool)
	go GetAsyncSolution(language, uuid, success)
	json.NewEncoder(w).Encode(map[string]bool{
		"result": <-success,
	})
}

func GetAsyncSolution(language string, uuid string, success chan bool) {
	success <- CheckSolution(language, uuid)
}

func CheckSolution(language string, uuid string) bool {
	err := exec.Command("docker", "run", language+"_test", uuid).Run()
	fmt.Println(err)
	return err == nil
}
