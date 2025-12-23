package usecase

import (
	"authentication/domain"
	"errors"
)

type ValidateToken struct {
}

func NewValidateTokenUser() *ValidateToken {
	return &ValidateToken{}
}

func (uc ValidateToken) Execute(tokenStr string) (string, error) {
	token := domain.Token(tokenStr)

	if token.IsExpired() {
		return "", errors.New("token expired")
	}

	if token.IsRevoked() {
		return "", errors.New("token revoked")
	}

	return token.UserId(), nil
}
