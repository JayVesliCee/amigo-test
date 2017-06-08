package main

import (
	"net/http"

	"github.com/pressly/chi"
)

func routes(r *chi.Mux) {
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Amigo test challenge API"))
	})

	r.Route("/messages", func(r chi.Router) {
		r.Post("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Messages route POST"))
		})
		r.Get("/:id", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Messages route Get"))
		})
	})
}
