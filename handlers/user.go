package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/beevk/go-todo/domain"
)

func (s *Server) registerUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		payload := domain.RegisterPayload{}

		err := json.NewDecoder(r.Body).Decode(&payload)
		if err != nil {
			badRequestResponse(w, err)
			return
		}

		fmt.Println("payload::", payload)

		// Validate payload

		// Store user in the database
		//user, err := s.domain.Register(payload)
		//if err != nil {
		//	http.Error(w, "Failed to register user", http.StatusInternalServerError)
		//	return
		//}
		// Implementation will go here
	}
}
