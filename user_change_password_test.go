package authentication

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestChangePassword_Valid(t *testing.T) {
	token := "token"
	oldPassword := "oldPassword"
	newPassword := "newPassword"
	err := ChangePassword(token, oldPassword, newPassword)
	assert.Equal(t, nil, err)
}
