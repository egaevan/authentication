package authentication

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRegistration_Valid(t *testing.T) {
	email := "testing@gmail.com"
	password := "testing123"
	user, err := Register(email, password)
	assert.Equal(t, nil, err)
	assert.NotEqual(t, password, user.password)
	assert.Equal(t, "active", user.status)
	assert.True(t, user.IsActive())
}

func TestRegistration_DuplicateEmail(t *testing.T) {
	email := "testing1@gmail.com"
	password := "testing123"
	user, err := Register(email, password)
	assert.Equal(t, errors.New("email already exists"), err)
	assert.Nil(t, user)
}

func TestRegistration_InvalidEmail(t *testing.T) {
	email := "testing1@gmail"
	password := "testing123"
	user, err := Register(email, password)
	assert.Equal(t, errors.New("invalid email format"), err)
	assert.Nil(t, user)
}

func TestRegistration_WeakPassword(t *testing.T) {
	email := "testing@gmail.com"
	password := "12345678"
	user, err := Register(email, password)
	assert.Equal(t, errors.New("password too weak"), err)
	assert.Nil(t, user)
}
