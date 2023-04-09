package types

type Stat struct{
	RPS float32 `json:"current_rps"`
	MedianResponseType float32 `json:"median_response_time"`
}

type StatsResponse struct{
	Stats []Stat 
}