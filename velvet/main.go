package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/J-Sumer/AutoScaler/velvet/docker"

	// "strconv"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	echo "github.com/labstack/echo/v4"
)

var URL = "http://152.7.179.7:8086"
var TOKEN = "TOKEN"
var ORG = "NCSU"
var BUCKET = "ADS"

func AddMetricEntry(cpu int, memory int) string{
    // Create a new client using an InfluxDB server base URL and an authentication token
	fmt.Println("Creating URL")
    client := influxdb2.NewClient(URL, TOKEN)
	fmt.Println("Creating WriteAPI")
    // Use blocking write client for writes to desired bucket
    writeAPI := client.WriteAPIBlocking(ORG, BUCKET)
    // Create point using full params constructor
    p := influxdb2.NewPoint("metric",
	map[string]string{"type": "stats"},
	map[string]interface{}{"cpu": cpu, "mem": memory},
	time.Now())
    // write point immediately
	fmt.Println("Writing started")
    writeAPI.WritePoint(context.Background(), p)
	fmt.Println("Writing completed")

    // Ensures background processes finishes
    client.Close()
	return "Added Entry in DB"
}


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

func addMetricEntryRoute(c echo.Context) error {
	cpu, _ := strconv.Atoi(c.Param("cpu"))
	memory, _ := strconv.Atoi(c.Param("memory"))
	return c.String(http.StatusOK, AddMetricEntry(cpu, memory))
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
	docker.CollectMetric()
	e := echo.New()
	e.GET("/", helloWorldRoute)
	e.GET("/create/container/:port/:name", createContainerRoute)
	e.GET("/delete/container/:id", deleteContainerRoute)
	e.GET("/count/containers", runningContainersCountRoute)
	e.GET("/metrics", metricsRoute)
	e.GET("/addMetric/:cpu/:memory", addMetricEntryRoute)


	// Run a go routine that collect metrics of the containers every x seconds

	e.Logger.Fatal(e.Start(":8000"))
}