package authentication

import "errors"

func ValidateToken(token string) (string, error) {
	if token == "token-expired" {
		return "", errors.New("token expired")
	}

	if token == "token-revoked" {
		return "", errors.New("token revoked")
	}

	return "user-id", nil
}
