package domain

import "fmt"

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

	fmt.Println(":I am being called", v.errors)

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
		Password: password,
	}

	user, err := d.DB.UserRepo.Create(data)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (d *Domain) hashPassword(password string) (string, error) {
	return password, nil
}
