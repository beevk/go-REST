package handlers

import (
	"net/http"

	"github.com/beevk/go-todo/utils"
	"github.com/go-chi/chi/v5"
)

func (s *Server) SetupRoutes(r *chi.Mux) {
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte("OK"))
		if err != nil {
			http.Error(w, "Failed to write response", http.StatusInternalServerError)
			return
		}
	})

	r.Route("/v1", func(r chi.Router) {
		r.Route("/users", func(r chi.Router) {
			r.Post("/register", s.registerUser())
		})

		r.Route("/todo", func(r chi.Router) {
			r.Use(s.withUser)
			r.Get("/", func(w http.ResponseWriter, r *http.Request) {
				response := map[string]string{"message": "GET /todo endpoint"}
				utils.JsonResponse(w, response, http.StatusOK)
			})
		})
	})
}
