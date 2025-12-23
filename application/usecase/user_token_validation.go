package usecase

import "errors"

type ValidateToken struct {
}

func NewValidateTokenUser() *ValidateToken {
	return &ValidateToken{}
}

func (uc ValidateToken) Execute(token string) (string, error) {
	if token == "token-expired" {
		return "", errors.New("token expired")
	}

	if token == "token-revoked" {
		return "", errors.New("token revoked")
	}

	return "user-id", nil
}
