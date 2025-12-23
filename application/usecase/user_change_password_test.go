package usecase_test

import (
	"authentication/application/usecase"
	"authentication/domain"
	"authentication/testdata"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestChangePassword_Valid(t *testing.T) {
	token := string(testdata.TokenValid)
	oldPassword := testdata.UserActive.Password()
	newPassword := "newPassword"
	changePw := usecase.NewChangePasswordUser(testdata.Users)
	err := changePw.Execute(token, string(oldPassword), newPassword)
	assert.Equal(t, nil, err)
}

func TestChangePassword_InvalidToken(t *testing.T) {
	token := "token"
	oldPassword := "oldPassword"
	newPassword := "newPassword"
	changePw := usecase.NewChangePasswordUser(testdata.Users)
	err := changePw.Execute(token, oldPassword, newPassword)
	assert.Equal(t, errors.New("invalid token"), err)
}

func TestChangePassword_UserNotFound(t *testing.T) {
	email, _ := domain.NewEmail("testing00@gmail.com")
	password, _ := domain.NewPassword("Testing123*")

	userNotFound := domain.NewUser(email, password, "active")
	token := domain.NewToken(userNotFound.IdString(), userNotFound.Email(), domain.TokenValid, "")

	oldPassword := "oldPassword"
	newPassword := "newPassword"
	changePw := usecase.NewChangePasswordUser(testdata.Users)
	err := changePw.Execute(string(token), oldPassword, newPassword)
	assert.Equal(t, errors.New("user not found"), err)
}

func TestChangePassword_InvalidOldPassword(t *testing.T) {
	token := string(testdata.TokenValid)
	oldPassword := "oldPassword-wrong"
	newPassword := "newPassword"
	changePw := usecase.NewChangePasswordUser(testdata.Users)
	err := changePw.Execute(token, oldPassword, newPassword)
	assert.Equal(t, errors.New("invalid old password"), err)
}
