package handlers

import (
	"net/http"

	"github.com/beevk/go-todo/domain"
	"github.com/beevk/go-todo/utils"
)

type authResponse struct {
	User  *domain.User     `json:"user"`
	Token *domain.JWTToken `json:"token"`
}

func (s *Server) registerUser() http.HandlerFunc {
	var payload domain.RegisterPayload

	return validatePayload(func(w http.ResponseWriter, r *http.Request) {
		user, err := s.domain.Register(payload)
		if err != nil {
			badRequestResponse(w, err)
			return
		}

		// Generate JWT token
		token, err := user.GenerateToken()
		if err != nil {
			internalServerErrorResponse(w, err)
			return
		}

		utils.JsonResponse(w, &authResponse{
			User:  user,
			Token: token,
		}, http.StatusCreated)
	}, &payload)
}
