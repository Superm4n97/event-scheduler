package server

import (
	"github.com/Superm4n97/event-scheduler/apis"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

var R = chi.NewRouter()

func pong(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("pong"))
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
}

func RouterSetup() {
	R.Use(middleware.Logger)

	R.Get("/ping", pong)
	R.Route("/apis", func(r chi.Router) {
		r.Use(apis.Authentication)
	})
}
