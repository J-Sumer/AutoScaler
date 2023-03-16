package main

import (
    // "fmt"
    "net/http"
	// "sync"

    // "autoScaler/influxDB"
    // "autoScaler/loadBalancer"

    "github.com/labstack/echo/v4"
)


func main() {
    // Adding entry to InfluxDB
    // fmt.Println("Adding metric details ")
    // influxDB.AddMetricEntry(200, 50)

    // Starting load balancer
    // loadBalancer.TestLoadBalancer()

    e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}