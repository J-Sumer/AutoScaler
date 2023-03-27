package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
	"time"
	"reflect"
)

func createContainer(port string, containerName string) string {
	portBind := port + ":8000"
	cmd := exec.Command("docker", "run", "--rm", "-p", portBind, "-d", containerName)
	var out strings.Builder
	cmd.Stdout = &out
	fmt.Println("Creating container\n")
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Container created with ID", out.String())
	return out.String()
}

func stopContainer(contId string) string {
	fmt.Println("Start: container stop: \n")
	fmt.Println(contId)
	cmdStop := exec.Command("docker", "stop", contId)
	var outStop strings.Builder
	cmdStop.Stdout = &outStop
	fmt.Println("Stopping container\n")
	err := cmdStop.Run()
	if err != nil {
		log.Fatal(err)
	}
	return out.String()
}

func main() {

	// Create container
	id := createContainer("8001", "node-app")

	time.Sleep(1 * time.Second)
	fmt.Println("Trying to stop container: " + id)
	fmt.Println("", reflect.TypeOf(id))
	// stopContainer(""+id)
	// stopContainer("b5ecdd2605b2772f454632524c75d2610471d4c7b286e0be920bc6d87a975104")

	// fmt.Println(id)
}