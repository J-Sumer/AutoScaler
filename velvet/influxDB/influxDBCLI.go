package influxdb

import (
	"context"
	"fmt"
	"time"

	stypes "github.com/J-Sumer/AutoScaler/velvet/docker/types"
	"github.com/J-Sumer/AutoScaler/velvet/locust"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

var URL = "http://152.7.179.7:8086"
var TOKEN = "TOKEN"
var ORG = "NCSU"
var BUCKET = "TrainingData"

func AddMetricEntry(cpu int, memory int, AllMetrics stypes.ContainerStats) string{
	// Get Metrics from locust
	RPS, MRT := locust.GetLocustMetrics()

    // Create a new client using an InfluxDB server base URL and an authentication token
	// fmt.Println("Creating URL")
    client := influxdb2.NewClientWithOptions(URL, TOKEN, influxdb2.DefaultOptions().SetBatchSize(20))
	// fmt.Println("Creating WriteAPI")
    // Use blocking write client for writes to desired bucket
    // writeAPI := client.WriteAPI(ORG, BUCKET)
    writeAPI := client.WriteAPIBlocking(ORG, BUCKET)

	fmt.Println("Adding metrics to influxdb")

    // Create point using full params constructor
	tag := map[string]string{"type": "agrigate"}
	fields := map[string]interface{}{}
	fields["cpu"] = cpu
	fields["mem"] = memory
	fields["RPS"] = RPS
	fields["MRT"] = MRT

    p := influxdb2.NewPoint("metrics", tag, fields, time.Now())

    // write point immediately
    writeAPI.WritePoint(context.Background(), p)

	// Create point for each container
	for i :=0; i<len(AllMetrics.AllMetrics); i++ {
		containerMetricDetails := AllMetrics.AllMetrics[i]
		cTag := map[string]string{}
		cTag["type"] = "container"
		cTag["id"] =  containerMetricDetails.ContainerId

		cFields := map[string]interface{}{}
		cFields["cpuUsage"] = containerMetricDetails.CPU * 5
		cFields["memUsage"] = containerMetricDetails.Memory
		// fmt.Println("ID", containerMetricDetails.ContainerId)
		// fmt.Println("CPU", containerMetricDetails.CPU)
		// fmt.Println("Memory", containerMetricDetails.Memory)
		p := influxdb2.NewPoint("metrics", cTag, cFields, time.Now())
		// writeAPI.WritePoint(p)
	    writeAPI.WritePoint(context.Background(), p)
	}

    client.Close()
	return "Added Entry in DB"
}