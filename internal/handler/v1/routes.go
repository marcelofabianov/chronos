package v1

import (
	"github.com/go-chi/chi/v5"
)

func LoadRoutes() *chi.Mux {
	r := chi.NewRouter()

	r.Route("/users", func(r chi.Router) {
		r.Get("/", GetUsersHandler)

		//...
	})

	return r
}
