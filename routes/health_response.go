package routes

type HealthResponse struct {
	Uptime       UptimeResponse       `json:"uptime"`
	HttpRequests HttpRequestsResponse `json:"httpRequests"`
}

type UptimeResponse struct {
	Value int64  `json:"value"`
	Unit  string `json:"unit"`
}

type HttpRequestsResponse struct {
	Counters map[string]int64 `json:"counters"`
}
