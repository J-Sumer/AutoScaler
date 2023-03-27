package main

import (
	"fmt"
	// "log"
	"os/exec"
	"bytes"
	// "strings"
	// "net/http"
	// "strconv"

	// "github.com/labstack/echo/v4"
)


func runningContainersCount() string {
	// cmdCount := exec.Command("docker", "ps", "-q", "|", "wc", "-l")
	// var outCount strings.Builder
	// cmdCount.Stdout = &outCount
	// err := cmdCount.Run()
	// if err != nil {
	// 	log.Fatal(err)
	// 	// return "Failed to fetch containers"
	// }
	// return outCount.String() 

	// cmd := exec.Command("find", "/", "-maxdepth", "1", "-exec", "wc", "-c", "{}", "\\")
	cmd := exec.Command("/bin/sh", "-c", "docker ps -q | wc -l")
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		return ""
	}
	fmt.Println("Result: " + out.String())
	return ""
}

func main() {

	fmt.Println(runningContainersCount())
}