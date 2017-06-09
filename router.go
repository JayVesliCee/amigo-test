package main

import "github.com/pressly/chi"

func routes(r *chi.Mux) {
	r.Route("/messages", func(r chi.Router) {
		r.Post("/", receiveMessageHandler)
		r.Get("/:id", retrieveMessageHandler)
	})
}
