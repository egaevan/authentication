package usecase_test

import (
	"authentication/application/usecase"
	"authentication/testdata"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLogout_Valid(t *testing.T) {
	token := string(testdata.TokenValid)
	logout := usecase.NewLogoutUser()
	err := logout.Execute(token)
	assert.Equal(t, nil, err)
}

func TestLogout_Invalid(t *testing.T) {
	token := "token-invalid"
	logout := usecase.NewLogoutUser()
	err := logout.Execute(token)
	assert.Equal(t, errors.New("invalid token"), err)
}
