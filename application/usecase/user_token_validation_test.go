package usecase_test

import (
	"authentication/application/usecase"
	"authentication/testdata"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTokenValidation_Valid(t *testing.T) {
	token := string(testdata.TokenValid)
	validate := usecase.NewValidateTokenUser()
	userId, err := validate.Execute(token)
	assert.Equal(t, nil, err)
	assert.NotEmpty(t, userId)
}

func TestTokenValidation_TokenExpired(t *testing.T) {
	token := string(testdata.TokenExpired)
	validate := usecase.NewValidateTokenUser()
	userId, err := validate.Execute(token)
	assert.Equal(t, errors.New("token expired"), err)
	assert.Empty(t, userId)
}

func TestTokenValidation_TokenRevoked(t *testing.T) {
	token := string(testdata.TokenRevoked)
	validate := usecase.NewValidateTokenUser()
	userId, err := validate.Execute(token)
	assert.Equal(t, errors.New("token revoked"), err)
	assert.Empty(t, userId)
}
