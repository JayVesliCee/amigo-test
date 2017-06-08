package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/pressly/chi"
	log "github.com/sirupsen/logrus"
)

/*
	TODO
	2 endpoints: post message + find message by id
	Postgresql in back
	Use external libs as less as possible
*/

func main() {
	pathToConfig := flag.String("config", "config.json", "Path to config file")
	flag.Parse()

	service, err := newService(*pathToConfig)
	if err != nil {
		log.WithError(err).Error("Error on initialising services")
	}

	r := chi.NewRouter()

	routes(r)

	fmt.Println("Running on port:", service.Config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", service.Config.Port), r))
}
