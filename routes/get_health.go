package routes

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/thenativeweb/monitoring/metrics"
)

const HealthName = "/health"

func GetHealth(collection *metrics.Collection) HttpHandler {
	collection.HttpRequests.RegisterRoute(HealthName)

	return func(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
		defer collection.HttpRequests.Increment(HealthName)

		uptime, uptimeUnit := collection.Uptime.Value()
		httpRequests := collection.HttpRequests.Value()

		healthResponse := HealthResponse{
			Uptime:       UptimeResponse{Value: uptime, Unit: uptimeUnit},
			HttpRequests: HttpRequestsResponse{Counters: httpRequests},
		}

		jsonResponse, err := json.Marshal(healthResponse)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)
		writer.Write(jsonResponse)
	}
}
