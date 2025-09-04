# Bug to fix

Route Registration (Happens Once): When your server starts, the setupUserRoutes method runs and registers routes:
`r.Post("/register", s.registerUser())`

This calls `registerUser()` immediately (not per request) and uses its return value as the handler.

Handler Creation (Happens Once): The `registerUser()` method executes and:
- Creates any local variables in its scope
- Returns a handler function via validatePayload
- The returned handler is stored by the router
- Request Handling (Happens Per Request):
- When requests arrive, the previously stored handler function is executed in a new goroutine.


The payload variable is:
- Created once when registerUser() runs during setup
- It Shared across all requests because:
  - It's in the closure's outer scope. The closure is created once but executed many times.
  - All executions reference the same memory location.

## Solution / Fix
The fix is moving the payload declaration inside the handler function so each request creates its own instance.\
OR\
Use middleware to bind and validate the payload.

```go
func (s *Server) registerUser() http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var payload domain.RegisterPayload
        if err := validateAndDecode(r, &payload); err != nil {
            badRequestResponse(w, err)
            return
        }
        
        user, err := s.domain.Register(payload)
        // rest of your handler...
    }
}

```

```go
// // PayloadValidation is the interface that all payload types must implement
type PayloadValidation interface {
IsValid() (bool, map[string]string)
}

// validateAndDecode decodes the request body into the provided payload and validates it
func validateAndDecode(r *http.Request, payload PayloadValidation) error {
    // Decode JSON from request body
    if err := json.NewDecoder(r.Body).Decode(payload); err != nil {
        return fmt.Errorf("invalid JSON: %w", err)
    }

    // Validate using the IsValid method from the interface
    hasErrors, errors := payload.IsValid()
    if hasErrors {
        // Convert the errors map to a structured error
        // You could return the first error, all errors, or create a custom error type
        for field, message := range errors {
            return fmt.Errorf("%s: %s", field, message)
        }
    }

    return nil
}
```
