package routes

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"github.com/J-Sumer/AutoScaler/velvet/types"
	"encoding/json"
)

func GetLocustMetrics() (float32, float32) {
	resp, err := http.Get("http://152.7.179.7:8089/stats/requests")
	var RPS float32 = 0.0
	var MRS float32 = 0.0
	if err != nil {
		fmt.Println("Error getting metrics from locust")
		return RPS, MRS
	}
	
	body1, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	var result types.StatsResponse
	json.Unmarshal(body1, &result)
	RPS = result.Stats[0].RPS
	MRS = result.Stats[0].MedianResponseType
	return RPS, MRS
}