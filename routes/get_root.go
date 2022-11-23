package routes

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/thenativeweb/monitoring/metrics"
)

const RootName = "/"

func GetRoot(message string, collection *metrics.Collection) HttpHandler {
	collection.HttpRequests.RegisterRoute(RootName)

	return func(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
		defer collection.HttpRequests.Increment(RootName)

		writer.WriteHeader(http.StatusOK)
		writer.Write([]byte(message + "\n"))
	}
}
