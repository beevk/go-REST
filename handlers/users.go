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

func (s *Server) getUserFromContext(r *http.Request) (*domain.User, error) {
	userCtx := r.Context().Value("currentUser")
	if userCtx == nil {
		return nil, domain.ErrNoResult
	}

	user, ok := userCtx.(*domain.User)
	if !ok {
		return nil, domain.ErrNoResult
	}

	return user, nil
}

func (s *Server) loginUser() http.HandlerFunc {
	var payload domain.LoginPayload
	return validatePayload(func(w http.ResponseWriter, r *http.Request) {
		user, err := s.domain.Login(payload)
		if err != nil {
			unauthorizedResponse(w, err)
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
		}, http.StatusOK)
	}, &payload)
}
