package locust

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"github.com/J-Sumer/AutoScaler/velvet/types"
	"encoding/json"
)

func GetLocustMetrics() (float32, float32) {
	resp, err := http.Get("http://152.7.178.162:8089/stats/requests")
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
	if len(result.Stats) == 1 {
		RPS = 0
		MRS = 0
	} else {
		RPS = ( result.Stats[1].RPS + result.Stats[2].RPS ) / 2
		MRS = (result.Stats[1].MedianResponseType + result.Stats[2].MedianResponseType) / 2
	}
	return RPS, MRS
}