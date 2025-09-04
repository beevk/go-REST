package handlers

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/beevk/go-todo/domain"
	"github.com/beevk/go-todo/utils"
	"github.com/go-chi/chi/v5"
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

//func (s *Server) getToDosByUserId() ([]*domain.ToDo, error) {
//	currentUser, err := s.getUserFromContext()
//	if err != nil {
//		unauthorizedResponse(w, err)
//		return nil, err
//	}
//
//	todo, err := s.domain.GetByUserId(currentUser.ID)
//	if err != nil {
//		internalServerErrorResponse(w, err)
//		return nil, err
//	}
//
//	utils.JsonResponse(w, todo, http.StatusOK)
//	return nil, nil
//}

func (s *Server) todoCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		todo := new(domain.ToDo)

		if todoId := chi.URLParam(r, "todoID"); todoId != "" {
			id, err := strconv.Atoi(todoId)
			if err != nil {
				badRequestResponse(w, err)
				return
			}

			todo, err = s.domain.Get(int64(id))
			if err != nil {
				response := map[string]string{"error": domain.ErrNoResult.Error()}
				utils.JsonResponse(w, response, http.StatusNotFound)
				return
			}
		}
		ctx := context.WithValue(r.Context(), "todo", todo)

		next.ServeHTTP(w, r.WithContext(ctx))
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

func (s *Server) getToDoFromContext(r *http.Request) (*domain.ToDo, error) {
	todoCtx := r.Context().Value("todo")
	if todoCtx == nil {
		return nil, domain.ErrNoResult
	}

	todo, ok := todoCtx.(*domain.ToDo)
	if !ok {
		return nil, domain.ErrNoResult
	}

	return todo, nil
}

func (s *Server) getToDoById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		todo, err := s.getToDoFromContext(r)
		if err != nil {
			if errors.Is(err, domain.ErrNoResult) {
				response := map[string]string{"error": domain.ErrNoResult.Error()}
				utils.JsonResponse(w, response, http.StatusNotFound)
				return
			}
			badRequestResponse(w, err)
			return
		}

		utils.JsonResponse(w, todo, http.StatusOK)
	}
}

func (s *Server) updateToDo() http.HandlerFunc {
	var updateToDoPayload domain.UpdateToDoPayload
	return validatePayload(func(w http.ResponseWriter, r *http.Request) {
		todo, err := s.getToDoFromContext(r)
		if err != nil {
			if errors.Is(err, domain.ErrNoResult) {
				response := map[string]string{"error": domain.ErrNoResult.Error()}
				utils.JsonResponse(w, response, http.StatusNotFound)
				return
			}
			badRequestResponse(w, err)
			return
		}

		todo, err = s.domain.Update(todo, &updateToDoPayload)
		if err != nil {
			internalServerErrorResponse(w, err)
			return
		}

		utils.JsonResponse(w, todo, http.StatusOK)
	}, &updateToDoPayload)
}

func (s *Server) deleteToDo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		todo, err := s.getToDoFromContext(r)
		if err != nil {
			if errors.Is(err, domain.ErrNoResult) {
				response := map[string]string{"error": domain.ErrNoResult.Error()}
				utils.JsonResponse(w, response, http.StatusNotFound)
				return
			}
			badRequestResponse(w, err)
			return
		}

		err = s.domain.Delete(todo)
		if err != nil {
			internalServerErrorResponse(w, err)
			return
		}
		w.WriteHeader(http.StatusOK)
		return
	}
}
