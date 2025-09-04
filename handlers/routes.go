package handlers

import (
	"net/http"

	"github.com/beevk/go-todo/utils"
	"github.com/go-chi/chi/v5"
)

func (s *Server) SetupRoutes(r *chi.Mux) {
	r.Get("/health", func(w http.ResponseWriter, _ *http.Request) {
		utils.JsonResponse(w, map[string]string{"status": "ok"}, http.StatusOK)
		return
	})

	r.Head("/health", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		return
	})

	r.Route("/v1", func(r chi.Router) {
		s.setupUserRoutes(r)
		s.setupTodoRoutes(r)
	})
}
