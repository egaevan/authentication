package authentication

import (
	"errors"
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

func TestChangePassword_InvalidOldPassword(t *testing.T) {
	token := "token"
	oldPassword := "oldPassword-wrong"
	newPassword := "newPassword"
	err := ChangePassword(token, oldPassword, newPassword)
	assert.Equal(t, errors.New("invalid old password"), err)
}
