package handlers

import (
	"fmt"
	"net/http"

	"github.com/beevk/go-todo/domain"
)

func (s *Server) registerUser() http.HandlerFunc {
	var payload domain.RegisterPayload

	return validatePayload(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("payload::", payload)
		fmt.Println("I AM HERE::")

		// Store user in the database
		//user, err := s.domain.Register(payload)
		//if err != nil {
		//	http.Error(w, "Failed to register user", http.StatusInternalServerError)
		//	return
		//}
		// Implementation will go here
	}, &payload)
}
