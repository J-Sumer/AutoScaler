package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
	"net/http"
	"time"
	// "strconv"

	"github.com/labstack/echo/v4"
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

func deleteContainer(contId string) string {
	fmt.Println("Start: container stop: \n")
	fmt.Println(contId)
	cmdStop := exec.Command("docker", "stop", contId)
	var outStop strings.Builder
	cmdStop.Stdout = &outStop
	fmt.Println("Stopping container\n")
	err := cmdStop.Run()
	if err != nil {
		return "Failed to stop container"
	}
	return outStop.String()
}

func runningContainersCount() string {
	cmdCount := exec.Command("/bin/sh", "-c", "docker ps -q | wc -l")
	var outCount strings.Builder
	cmdCount.Stdout = &outCount
	err := cmdCount.Run()
	if err != nil {
		return "Failed to fetch containers"
	}
	
	return outCount.String() 
}

func getCPUMetric() string {
	start := time.Now()
	cmdCPU := exec.Command("/bin/sh", "-c", "./statCollect.sh")
	// cmdCPU := exec.Command("/bin/sh", "-c", "free -m | awk 'NR==2{printf "%.2f%%", $3*100/$2 }'")
	var outCPU strings.Builder
	cmdCPU.Stdout = &outCPU
	err := cmdCPU.Run()
	if err != nil {
		return "Failed to fetch CPU metrics"
	}
	duration := time.Since(start)
	fmt.Println("Time to get stats: ", duration)
	return outCPU.String()
}

func runningContainersCountRoute(c echo.Context) error {
	return c.String(http.StatusOK, runningContainersCount())
}

func createContainerRoute(c echo.Context) error {
  	// User ID from path `users/:id`
  	port := c.Param("port")
  	name := c.Param("name")
	return c.String(http.StatusOK, createContainer(port, name))
}

func deleteContainerRoute(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, deleteContainer(id))
}

func helloWorldRoute(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World")
}



func metricsRoute(c echo.Context) error {
	return c.String(http.StatusOK, getCPUMetric())
}

func Notes() {
	// To count the number of containers running 
	//docker ps -q | wc -l
	//docker ps | grep imagename | wc -l
	//docker inspect --format='{{.Config.Image}}' $(docker ps -q) | grep imagename | wc -l
	
	// To convert num to string
	// count, err := strconv.Atoi(outCount.String())
}

func main() {

	e := echo.New()
	e.GET("/", helloWorldRoute)
	e.GET("/create/container/:port/:name", createContainerRoute)
	e.GET("/delete/container/:id", deleteContainerRoute)
	e.GET("/count/containers", runningContainersCountRoute)
	e.GET("/metrics", metricsRoute)

	e.Logger.Fatal(e.Start(":8000"))
}