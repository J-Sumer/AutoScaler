package main

import (
    "fmt"
    "influxdb/influxDB"
)

func main() {
    fmt.Println("Adding metric details ")
    influxDB.AddMetricEntry(200, 50)
}