package handlers

import (
	"time"

	"github.com/beevk/go-todo/domain"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Server struct {
	domain *domain.Domain
}

func NewServer(d *domain.Domain) *Server {
	return &Server{domain: d}
}

func setupMiddleware(r *chi.Mux) {
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Compress(6, "application/json"))
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(middleware.Timeout(60 * time.Second))
}

func (s *Server) SetupRouter() *chi.Mux {
	r := chi.NewRouter()

	setupMiddleware(r)

	s.SetupRoutes(r)

	return r
}
