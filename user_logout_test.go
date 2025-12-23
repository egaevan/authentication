package authentication

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLogout_Valid(t *testing.T) {
	token := "token"
	err := Logout(token)
	assert.Equal(t, nil, err)
}
