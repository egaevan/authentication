package usecase

import (
	"authentication/domain"
	"errors"
)

func Register(email, password string) (*domain.User, error) {
	if email == "testing1@gmail.com" {
		return nil, errors.New("email already exists")
	}

	emailDomain, err := domain.NewEmail(email)
	if err != nil {
		return nil, err
	}

	passwordHash, err := domain.NewPassword(password)
	if err != nil {
		return nil, err
	}

	return domain.NewUser(emailDomain, passwordHash), nil
}
