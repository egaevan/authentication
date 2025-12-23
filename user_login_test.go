package authentication

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLogin_Valid(t *testing.T) {
	email := "testing@gmail.com"
	password := "Testing123*"

	token, err := Login(email, password)
	assert.Equal(t, nil, err)
	assert.NotEmpty(t, token)
}

func TestLogin_InvalidPassword(t *testing.T) {
	email := "testing@gmail.com"
	password := "Testing123#"

	token, err := Login(email, password)
	assert.Equal(t, errors.New("invalid credentials"), err)
	assert.Empty(t, token)
}
