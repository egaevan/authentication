package usecase

import (
	"authentication/domain"
	"errors"
)

type Register struct {
	users map[string]*domain.User
}

func NewRegisterUser(users map[string]*domain.User) *Register {
	return &Register{users: users}
}

func (uc Register) Execute(email, password string) (*domain.User, error) {
	_, ok := uc.users[email]
	if ok {
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

	return domain.NewUser(emailDomain, passwordHash, ""), nil
}
