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
		return nil, ErrorUserWithEmailAlreadyExist
	}

	userExists, _ = d.DB.UserRepo.GetByEmail(payload.Username)
	if userExists != nil {
		return nil, ErrorUserWithUsernameAlreadyExist
	}

	password, err := d.setPassword(payload.Password)
	if err != nil {
		return nil, err
	}

	data := &User{
		Username: payload.Username,
		Password: *password,
		Email:    payload.Email,
	}

	user, err := d.DB.UserRepo.Create(data)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (d *Domain) setPassword(password string) (*string, error) {

	return nil, nil
}
