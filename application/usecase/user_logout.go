package usecase

import (
	"errors"
)

type Logout struct {
}

func NewLogoutUser() *Logout {
	return &Logout{}
}

func (uc Logout) Execute(token string) error {
	if token != "token" {
		return errors.New("invalid token")
	}

	return nil
}
