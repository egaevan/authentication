package authentication

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLogout_Valid(t *testing.T) {
	token := "token"
	err := Logout(token)
	assert.Equal(t, nil, err)
}

func TestLogout_Invalid(t *testing.T) {
	token := "token-invalid"
	err := Logout(token)
	assert.Equal(t, errors.New("invalid token"), err)
}
