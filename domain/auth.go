package domain

type RegisterPayload struct {
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
	Username        string `json:"username"`
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
