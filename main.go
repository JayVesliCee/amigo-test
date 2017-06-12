/*
	The main application is using a lightway framework (chi) in order to use a optimised and stable router.
*/

package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/pressly/chi"
	"github.com/pressly/chi/middleware"
	log "github.com/sirupsen/logrus"
)

func main() {
	pathToConfig := flag.String("config", "config.json", "Path to config file")
	flag.Parse()

	service, err := newService(*pathToConfig)
	if err != nil {
		log.WithError(err).Error("Error on initialising services")
	}
	defer service.Close()

	r := chi.NewRouter()

	r.Use(middleware.StripSlashes)

	routes(r, service)

	log.Info("Running on port:", service.Config.Port)
	log.WithError(http.ListenAndServe(fmt.Sprintf(":%d", service.Config.Port), r))
}
