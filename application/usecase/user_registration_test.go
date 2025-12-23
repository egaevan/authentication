package usecase_test

import (
	"authentication/application/usecase"
	"authentication/testdata"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRegistration_Valid(t *testing.T) {
	email := "testing1@gmail.com"
	password := "Testing123*"

	register := usecase.NewRegisterUser(testdata.Users)
	user, err := register.Execute(email, password)
	assert.Equal(t, nil, err)
	assert.NotEqual(t, password, user.Password())
	assert.Equal(t, "active", user.Status())
	assert.True(t, user.IsActive())
}

func TestRegistration_DuplicateEmail(t *testing.T) {
	email := "testing@gmail.com"
	password := "testing123"
	register := usecase.NewRegisterUser(testdata.Users)
	user, err := register.Execute(email, password)
	assert.Equal(t, errors.New("email already exists"), err)
	assert.Nil(t, user)
}

func TestRegistration_InvalidEmail(t *testing.T) {
	email := "testing1@gmail"
	password := "testing123"
	register := usecase.NewRegisterUser(testdata.Users)
	user, err := register.Execute(email, password)
	assert.Equal(t, errors.New("invalid email format"), err)
	assert.Nil(t, user)
}

func TestRegistration_WeakPassword(t *testing.T) {
	email := "testing1@gmail.com"
	password := "12345678"
	register := usecase.NewRegisterUser(testdata.Users)
	user, err := register.Execute(email, password)
	assert.Equal(t, errors.New("password too weak"), err)
	assert.Nil(t, user)
}
