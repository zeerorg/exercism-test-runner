package main

import (
	"fmt"
	"os/exec"
)

func GetAsyncSolution(language string, uuid string, success chan bool) {
	success <- CheckSolution(language, uuid)
}

func CheckSolution(language string, uuid string) bool {
	err := exec.Command("docker", "run", language+"_test", uuid).Run()
	fmt.Println(err)
	return err == nil
}
