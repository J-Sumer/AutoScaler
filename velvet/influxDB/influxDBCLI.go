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
// var BUCKET = "ADS"
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

    // Create point using full params constructor
	tag := map[string]string{"type": "agrigate"}
	fields := map[string]interface{}{}
	fields["cpu"] = cpu
	fields["mem"] = memory
	fields["RPS"] = RPS
	fields["MRT"] = MRT

    p := influxdb2.NewPoint("metrics", tag, fields, time.Now())

    // write point immediately
	// fmt.Println("Writing started")
    // writeAPI.WritePoint(p)
    writeAPI.WritePoint(context.Background(), p)
	// fmt.Println("Writing completed")

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

	fmt.Println("Adding metrics to influxdb")

	// testTag := map[string]string{}
	// testTag["tag1"] = "testtag1"
	// testTag["tag2"] = "testtag2"

	// testField := map[string]interface{}{}
	// testField["field1"] = 10
	// testField["field2"] = 30

	// p = influxdb2.NewPoint("metric", testTag, testField, time.Now())

	// containerMetricDetails := AllMetrics.AllMetrics[0]
	// containerTag := map[string]string{}
	// containerTag["type"] = "containertest"
	// containerTag["id"] =  containerMetricDetails.ContainerId

	// // containerField := map[string]interface{}{}
	// // containerField["cpu"] = containerMetricDetails.CPU
	// // containerField["mem"] = containerMetricDetails.Memory

	// testField := map[string]interface{}{}
	// testField["cpu"] = 10
	// testField["mem"] = 30

	// fmt.Println("Adding for container", containerMetricDetails.ContainerId)
	// fmt.Println("CPU", containerMetricDetails.CPU)

	// p = influxdb2.NewPoint("metric", containerTag, testField, time.Now())

	
	// fields1 := map[string]interface{}{}
	// fields1["cpu"] = cpu
	// fields1["mem"] = memory
	// fields1["RPS"] = RPS
	// fields1["test"] = 10
	// fields1["MRT"] = MRT

	// fields1 := map[string]interface{}{}
	// fields1["cpu"] = containerMetricDetails.CPU
	// fields1["mem"] = containerMetricDetails.Memory
	// fields1["testf"] = 100
	// p = influxdb2.NewPoint("metric", containerTag, fields1, time.Now())
	// writeAPI.WritePoint(p)
	// writeAPI.WritePoint(context.Background(), p)

	// Force all unwritten data to be sent
    // writeAPI.Flush()
    // Ensures background processes finishes
    client.Close()
	return "Added Entry in DB"
}