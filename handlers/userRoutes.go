package handlers

import (
	"github.com/go-chi/chi/v5"
)

func (s *Server) setupUserRoutes(r chi.Router) {
	r.Route("/users", func(r chi.Router) {
		r.Post("/register", s.registerUser())
	})
}
