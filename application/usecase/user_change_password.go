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

func (uc ChangePassword) Execute(tokenStr, oldPassword, newPassword string) error {
	token := domain.Token(tokenStr)

	if !token.IsValid() {
		return errors.New("invalid token")
	}

	user, ok := uc.users[token.Email()]
	if !ok {
		return errors.New("user not found")
	}

	if user.InvalidPassword(oldPassword) {
		return errors.New("invalid old password")
	}

	return nil
}
