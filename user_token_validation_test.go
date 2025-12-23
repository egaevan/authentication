package authentication

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTokenValidation_Valid(t *testing.T) {
	token := "token"
	userId, err := ValidateToken(token)
	assert.Equal(t, nil, err)
	assert.NotEmpty(t, userId)
}

func TestTokenValidation_TokenExpired(t *testing.T) {
	token := "token-expired"
	userId, err := ValidateToken(token)
	assert.Equal(t, errors.New("token expired"), err)
	assert.Empty(t, userId)
}

func TestTokenValidation_TokenRevoked(t *testing.T) {
	token := "token-revoked"
	userId, err := ValidateToken(token)
	assert.Equal(t, errors.New("token revoked"), err)
	assert.Empty(t, userId)
}
