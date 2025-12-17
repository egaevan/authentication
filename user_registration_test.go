package authentication

import (
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
