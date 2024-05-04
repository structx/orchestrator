package router

import "github.com/go-chi/chi/v5"

// New
func New() *chi.Mux {

	r := chi.NewRouter()

	return r
}
