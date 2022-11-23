package main

import (
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/zerolog/log"
	"github.com/thenativeweb/monitoring/metrics"
	"github.com/thenativeweb/monitoring/routes"
)

func main() {
	collection := metrics.NewCollection(time.Now())

	router := httprouter.New()

	router.GET(routes.RootName, routes.GetRoot("Hello, world!", collection))
	router.GET(routes.HealthName, routes.GetHealth(collection))

	server := http.Server{Addr: ":3000", Handler: router}
	err := server.ListenAndServe()

	log.Fatal().Err(err).Msg("server failed")
}
