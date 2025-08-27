package handlers

import (
	"encoding/json"
	"net/http"
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

func SetupRouter(d *domain.Domain) *chi.Mux {
	// Setup your routes here
	server := NewServer(d)
	r := chi.NewRouter()

	setupMiddleware(r)

	server.SetupRoutes(r)

	return r
}

func jsonResponse(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if data == nil {
		data = map[string]string{}
	}

	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	return
}

func badRequestResponse(w http.ResponseWriter, err error) {
	data := map[string]string{"error": err.Error()}
	jsonResponse(w, data, http.StatusBadRequest)
}
