package main

import (
	"os/exec"
)

func GetAsyncSolution(language string, uuid string, success chan bool) {
	success <- CheckSolution(language, uuid)
}

func CheckSolution(language string, uuid string) bool {
	_, err := exec.Command("docker", "run", language+"_test", uuid).CombinedOutput()
	return err == nil
}
