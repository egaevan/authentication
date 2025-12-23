package domain

import "encoding/base64"

type Token string

const tokenSecret = "secret"

func NewToken(userId string, email Email, status TokenStatus) Token {
	if status == "" {
		status = TokenValid
	}

	raw := userId + string(email) + ":" + string(status) + ":" + tokenSecret
	return Token(base64.StdEncoding.EncodeToString([]byte(raw)))
}
