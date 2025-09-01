package domain

import (
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/golang-jwt/jwt/v5/request"
	"golang.org/x/crypto/bcrypt"
)

type RegisterPayload struct {
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
	Username        string `json:"username"`
}

func (r *RegisterPayload) IsValid() (bool, map[string]string) {
	v := NewValidator()

	v.MustNotBeEmpty("email", r.Email)
	v.MustBeLongerThan("email", r.Email, 6)
	v.MustBeValidEmail("email", r.Email)

	v.MustNotBeEmpty("username", r.Username)
	v.MustBeLongerThan("username", r.Username, 4)

	v.MustNotBeEmpty("password", r.Password)
	v.MustBeLongerThan("password", r.Password, 6)

	v.MustNotBeEmpty("confirmPassword", r.ConfirmPassword)
	v.MustMatch("password", r.Password, "confirmPassword", r.ConfirmPassword)

	return v.HasErrors(), v.errors
}

func (d *Domain) Register(payload RegisterPayload) (*User, error) {
	userExists, _ := d.DB.UserRepo.GetByEmail(payload.Email)
	if userExists != nil {
		return nil, ErrUserWithEmailAlreadyExists
	}

	userExists, _ = d.DB.UserRepo.GetByUsername(payload.Username)
	if userExists != nil {
		return nil, ErrUserWithUsernameAlreadyExists
	}

	password, err := d.hashPassword(payload.Password)
	if err != nil {
		return nil, err
	}

	data := &User{
		Email:    payload.Email,
		Username: payload.Username,
		Password: *password,
	}

	user, err := d.DB.UserRepo.Create(data)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (d *Domain) hashPassword(password string) (*string, error) {
	bytePassword := []byte(password)
	hashedPassword, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	hashedPasswordStr := string(hashedPassword)
	return &hashedPasswordStr, nil
}

func stripBearerPrefixFromToken(token string) (string, error) {
	const bearerPrefix = "Bearer "
	// Check if the token starts with "Bearer " and strip it
	if len(token) > len(bearerPrefix) && token[0:len(bearerPrefix)] == bearerPrefix {
		return token[len(bearerPrefix):], nil
	}
	return token, nil // Return whole string if the prefix ("Bearer") is not present
}

// Reads the "Authorization" header from the HTTP request
var authHeaderExtractor = &request.PostExtractionFilter{
	Extractor: request.HeaderExtractor{"Authorization", "authorization"},
	Filter:    stripBearerPrefixFromToken,
}

var authExtractor = &request.MultiExtractor{
	authHeaderExtractor,
}

func ParseToken(r *http.Request) (*jwt.Token, error) {
	token, err := request.ParseFromRequest(r, authExtractor, func(t *jwt.Token) (interface{}, error) {
		b := []byte(os.Getenv("JWT_SECRET"))
		return b, nil
	})
	return token, err
}
