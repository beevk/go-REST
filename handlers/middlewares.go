package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/beevk/go-todo/domain"
	. "github.com/beevk/go-todo/utils"
	"github.com/golang-jwt/jwt/v5"
)

func badRequestResponse(w http.ResponseWriter, err error) {
	data := map[string]string{"error": err.Error()}
	JsonResponse(w, data, http.StatusBadRequest)
	return
}

func unauthorizedResponse(w http.ResponseWriter, err error) {
	data := map[string]string{"error": "unauthorized"}
	if err != nil {
		data["error"] = err.Error()
	}
	JsonResponse(w, data, http.StatusUnauthorized)
	return
}

func forbiddenResponse(w http.ResponseWriter, err error) {
	data := map[string]string{"error": "forbidden"}
	if err != nil {
		data["error"] = err.Error()
	}
	JsonResponse(w, data, http.StatusForbidden)
	return
}

func internalServerErrorResponse(w http.ResponseWriter, err error) {
	data := map[string]string{"error": "internal server error"}
	fmt.Println("Internal Server Error:", err)
	JsonResponse(w, data, http.StatusInternalServerError)
	return
}

// Middleware that validates the request payload
func validatePayload(next http.HandlerFunc, payload domain.PayloadValidation) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := json.NewDecoder(r.Body).Decode(&payload)
		if err != nil {
			badRequestResponse(w, err)
			return
		}

		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				fmt.Println("Error closing request body:", err)
			}
		}(r.Body)

		if isValid, errs := payload.IsValid(); !isValid {
			JsonResponse(w, errs, http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), "payload", payload)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

func (s *Server) withUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := domain.ParseToken(r)
		if err != nil {
			unauthorizedResponse(w, err)
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			userId := int64(claims["user_id"].(float64))

			user, err := s.domain.GetUserById(userId)
			if err != nil {
				unauthorizedResponse(w, err)
				return
			}

			ctx := context.WithValue(r.Context(), "currentUser", user)
			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			unauthorizedResponse(w, nil)
			return
		}
	})
}

func (s *Server) validateOwnership(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		todo, err := s.getToDoFromContext(r)
		if err != nil {
			badRequestResponse(w, err)
			return
		}

		user, err := s.getUserFromContext(r)
		if err != nil {
			unauthorizedResponse(w, err)
			return
		}

		if todo.UserID != user.ID {
			unauthorizedResponse(w, errors.New("unauthorized"))
			return
		}

		next.ServeHTTP(w, r)
	})
}

// validateOwnership middleware is too specific to ToDo entity
// we can make it more generic by using interface
// any entity that has an owner should implement this interface
// and we can use this middleware to validate ownership of that entity
// Example: ToDo, Project, etc.
// This middleware will check if the current user is the owner of the entity
// The entity should be set in the context with the given contextKey
// Example usage: r.Context().Value("todo").(*domain.ToDo)
func (s *Server) withOwner(contextKey string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			user, err := s.getUserFromContext(r)
			if err != nil {
				unauthorizedResponse(w, err)
				return
			}

			isOwner := r.Context().Value(contextKey).(domain.HasOwner).IsOwner(user)
			if !isOwner {
				forbiddenResponse(w, domain.ErrForbidden)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
