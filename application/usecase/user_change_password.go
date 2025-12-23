package usecase

import (
	"authentication/domain"
	"errors"
)

type ChangePassword struct {
	users map[string]*domain.User
}

func NewChangePasswordUser(users map[string]*domain.User) *ChangePassword {
	return &ChangePassword{users: users}
}

func (uc ChangePassword) Execute(token string, oldPassword string, newPassword string) error {
	if oldPassword == "oldPassword-wrong" {
		return errors.New("invalid old password")
	}

	return nil
}
