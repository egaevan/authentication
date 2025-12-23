package usecase

import (
	"authentication/domain"
	"errors"
)

type Login struct {
	users map[string]*domain.User
}

func NewLoginUser(users map[string]*domain.User) *Login {
	return &Login{users: users}
}

func (uc Login) Execute(email string, password string) (token string, err error) {
	user, ok := uc.users[email]
	if !ok {
		return "", errors.New("invalid credentials")
	}

	if !user.IsActive() {
		return "", errors.New("user inactive")
	}

	passwordVO, err := domain.NewPassword(password)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	if passwordVO != user.Password() {
		return "", errors.New("invalid credentials")
	}

	tokenVO := domain.NewToken(user.IdString(), user.Email(), "", "")

	return string(tokenVO), nil
}
