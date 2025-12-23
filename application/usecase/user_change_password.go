package usecase

import "errors"

type ChangePassword struct {
}

func NewChangePasswordUser() *ChangePassword {
	return &ChangePassword{}
}

func (uc ChangePassword) Execute(token string, oldPassword string, newPassword string) error {
	if oldPassword == "oldPassword-wrong" {
		return errors.New("invalid old password")
	}

	return nil
}
