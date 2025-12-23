package usecase_test

import (
	"authentication/application/usecase"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestChangePassword_Valid(t *testing.T) {
	token := "token"
	oldPassword := "oldPassword"
	newPassword := "newPassword"
	changePw := usecase.NewChangePasswordUser()
	err := changePw.Execute(token, oldPassword, newPassword)
	assert.Equal(t, nil, err)
}

func TestChangePassword_InvalidOldPassword(t *testing.T) {
	token := "token"
	oldPassword := "oldPassword-wrong"
	newPassword := "newPassword"
	changePw := usecase.NewChangePasswordUser()
	err := changePw.Execute(token, oldPassword, newPassword)
	assert.Equal(t, errors.New("invalid old password"), err)
}
