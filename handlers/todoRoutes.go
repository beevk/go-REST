package handlers

import (
	"github.com/go-chi/chi/v5"
)

func (s *Server) setupTodoRoutes(r chi.Router) {
	r.Route("/todo", func(r chi.Router) {
		r.Use(s.withUser)

		r.Post("/", s.createToDo())

		//r.Get("/", s.getAllToDos())

		r.Route("/{todoID}", func(r chi.Router) {
			r.Use(s.todoCtx)
			r.Use(s.validateOwnership)

			r.Get("/", s.getToDoById())
			r.Patch("/", s.updateToDo())
			r.Delete("/", s.deleteToDo())
		})
	})
}
