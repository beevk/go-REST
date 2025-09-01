package domain

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type User struct {
	ID       int64  `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"-"` // Exclude password from JSON responses

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type JWTToken struct {
	AccessToken string      `json:"accessToken"`
	ExpiresAt   time.Time   `json:"expiresAt"`
	Payload     interface{} `json:"payload"`
}

func (u *User) GenerateToken() (*JWTToken, error) {
	jwtToken := jwt.New(jwt.GetSigningMethod("HS256"))

	expiresAt := time.Now().Add(time.Hour * 24 * 7) // Token valid for 7 days

	jwtToken.Claims = jwt.MapClaims{
		"user_id": u.ID,
		"exp":     expiresAt.Unix(),
	}

	accessToken, err := jwtToken.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return nil, err
	}

	return &JWTToken{
		AccessToken: accessToken,
		ExpiresAt:   expiresAt,
	}, nil
}

func (d *Domain) GetUserById(id int64) (*User, error) {
	user, err := d.DB.UserRepo.GetById(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}
