package main

import (
	"github.com/J-Sumer/AutoScaler/velvet/docker"
	"github.com/J-Sumer/AutoScaler/velvet/routes"

	echo "github.com/labstack/echo/v4"
)

func Notes() {
	// To count the number of containers running 
	//docker ps -q | wc -l
	//docker ps | grep imagename | wc -l
	//docker inspect --format='{{.Config.Image}}' $(docker ps -q) | grep imagename | wc -l
	
	// To convert num to string
	// count, err := strconv.Atoi(outCount.String())
}

func main() {
	go docker.MetricExporter(15)
	e := echo.New()
	e.GET("/", routes.HelloWorldRoute)
	e.GET("/container/create/:port/:name", routes.CreateContainerRoute)
	e.GET("/container/delete/:id", routes.DeleteContainerRoute)
	e.GET("/container/count", routes.RunningContainersCountRoute)
	e.GET("/metrics", routes.MetricsRoute)
	e.GET("/addMetric/:cpu/:memory", routes.AddMetricEntryRoute)

	// Run a go routine that collect metrics of the containers every x seconds

	e.Logger.Fatal(e.Start(":8000"))
}