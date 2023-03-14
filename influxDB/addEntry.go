package influxDB

import (
    "context"
    "fmt"
    "time"

    "github.com/influxdata/influxdb-client-go/v2"
)

var URL = "http://localhost:8086"
var TOKEN = "TOKEN"
var ORG = "NCSU"
var BUCKET = "Metrics"

// measurementName,tagKey=tagValue fieldKey="fieldValue" 1465839830100400200
// This is how data is stored in InfluxDB
func AddMetricEntry(cpu int, memory int) {
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

}


func AddEntryTemplate() {
    // Create a new client using an InfluxDB server base URL and an authentication token
    client := influxdb2.NewClient("http://localhost:8086", "TOKEN")
    // Use blocking write client for writes to desired bucket
    writeAPI := client.WriteAPIBlocking("NCSU", "bucket")
    // Create point using full params constructor
    p := influxdb2.NewPoint("stat",
        map[string]string{"unit": "temperature"},
        map[string]interface{}{"avg": 25.5, "max": 106.0},
        time.Now())
    // write point immediately
    writeAPI.WritePoint(context.Background(), p)
    // // Create point using fluent style
    // p = influxdb2.NewPointWithMeasurement("stat").
    //     AddTag("unit", "temperature").
    //     AddField("avg", 23.2).
    //     AddField("max", 65.0).
    //     SetTime(time.Now())
    // err := writeAPI.WritePoint(context.Background(), p)
	// if err != nil {
	// 	panic(err)
	// }
    // // Or write directly line protocol
    // line := fmt.Sprintf("stat,unit=temperature avg=%f,max=%f", 23.5, 55.0)
    // err = writeAPI.WriteRecord(context.Background(), line)
	// if err != nil {
	// 	panic(err)
	// }

    // Get query client
    queryAPI := client.QueryAPI("NCSU")
    // Get parser flux query result
    result, err := queryAPI.Query(context.Background(), `from(bucket:"bucket")|> range(start: -1h) |> filter(fn: (r) => r._measurement == "stat")`)
    if err == nil {
        // Use Next() to iterate over query result lines
        for result.Next() {
            // Observe when there is new grouping key producing new table
            if result.TableChanged() {
                fmt.Printf("table: %s\n", result.TableMetadata().String())
            }
            // read result
            fmt.Printf("row: %s\n", result.Record().String())
        }
        if result.Err() != nil {
            fmt.Printf("Query error: %s\n", result.Err().Error())
        }
    } else {
		panic(err)
    }
    // Ensures background processes finishes
    client.Close()

}