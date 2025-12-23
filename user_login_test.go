package authentication

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLogin_Valid(t *testing.T) {
	email := "testing@gmail.com"
	password := "Testing123*"

	token, err := Login(email, password)
	assert.Equal(t, nil, err)
	assert.NotNil(t, token)
}
