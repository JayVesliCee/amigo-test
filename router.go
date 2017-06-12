package main

import "github.com/pressly/chi"

func routes(r *chi.Mux, s *service) {
	r.Route("/messages", func(r chi.Router) {
		r.Post("/", s.receiveMessageHandler)
		r.Get("/:id", s.retrieveMessageHandler)
	})
}
