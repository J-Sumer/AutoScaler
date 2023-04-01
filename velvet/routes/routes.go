package routes

import (
	"context"
	"fmt"
	"log"
	"os/exec"
	"strings"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
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


func CreateContainer(port string, containerName string) string {
	portBind := port + ":8000"
	cmd := exec.Command("docker", "run", "--rm", "-p", portBind, "-d", containerName)
	var out strings.Builder
	cmd.Stdout = &out
	fmt.Println("Creating container")
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Container created with ID", out.String())
	return out.String()
}

func DeleteContainer(contId string) string {
	fmt.Println("Start: container stop:")
	fmt.Println(contId)
	cmdStop := exec.Command("docker", "stop", contId)
	var outStop strings.Builder
	cmdStop.Stdout = &outStop
	fmt.Println("Stopping container")
	err := cmdStop.Run()
	if err != nil {
		return "Failed to stop container"
	}
	return outStop.String()
}

func RunningContainersCount() string {
	cmdCount := exec.Command("/bin/sh", "-c", "docker ps -q | wc -l")
	var outCount strings.Builder
	cmdCount.Stdout = &outCount
	err := cmdCount.Run()
	if err != nil {
		return "Failed to fetch containers"
	}
	
	return outCount.String() 
}

func GetCPUMetric() string {
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