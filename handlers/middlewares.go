package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	. "github.com/beevk/go-todo/utils"
)

func badRequestResponse(w http.ResponseWriter, err error) {
	data := map[string]string{"error": err.Error()}
	JsonResponse(w, data, http.StatusBadRequest)
}

// Middleware that validates the request payload
func validatePayload(next http.HandlerFunc, payload PayloadValidation) http.HandlerFunc {
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
