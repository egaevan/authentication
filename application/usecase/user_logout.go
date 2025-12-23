package usecase

import (
	"authentication/domain"
	"errors"
)

type Logout struct {
}

func NewLogoutUser() *Logout {
	return &Logout{}
}

func (uc Logout) Execute(tokenStr string) error {
	token := domain.Token(tokenStr)

	if !token.IsValid() {
		return errors.New("invalid token")
	}

	return nil
}
