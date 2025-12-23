package authentication

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTokenValidation_Valid(t *testing.T) {
	token := "token"
	userId, err := ValidateToken(token)
	assert.Equal(t, nil, err)
	assert.NotEmpty(t, userId)
}
