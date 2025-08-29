package handlers

import (
	"net/http"

	"github.com/beevk/go-todo/domain"
	"github.com/beevk/go-todo/utils"
)

func (s *Server) registerUser() http.HandlerFunc {
	var payload domain.RegisterPayload

	return validatePayload(func(w http.ResponseWriter, r *http.Request) {
		user, err := s.domain.Register(payload)
		if err != nil {
			badRequestResponse(w, err)
			return
		}

		// Generate JWT token

		utils.JsonResponse(w, user, http.StatusCreated)
	}, &payload)
}
