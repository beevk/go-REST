package handlers

import (
	"net/http"

	"github.com/beevk/go-todo/domain"
	"github.com/beevk/go-todo/utils"
)

type todoResponse struct {
	Todo *domain.ToDo `json:"todo"`
}

func (s *Server) createToDo() http.HandlerFunc {
	var payload domain.CreateToDoPayload
	return validatePayload(func(w http.ResponseWriter, r *http.Request) {
		currentUser, err := s.getUserFromContext(r)
		if err != nil {
			unauthorizedResponse(w, err)
			return
		}

		todo, err := s.domain.Create(payload, currentUser)
		if err != nil {
			internalServerErrorResponse(w, err)
			return
		}

		//response := todoResponse{Todo: todo}

		utils.JsonResponse(w, todo, http.StatusCreated)
	}, &payload)
}
