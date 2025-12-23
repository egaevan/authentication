package domain

import (
	"encoding/base64"
	"strconv"
	"strings"
	"time"
)

type Token string

const tokenSecret = "secret"

func NewToken(userId string, email Email, status TokenStatus, timeUnix string) Token {
	if status == "" {
		status = TokenValid
	}
	if timeUnix == "" {
		timeUnix = strconv.FormatInt(time.Now().Unix(), 10)
	}

	raw := userId + ":" + string(email) + ":" + string(status) + ":" +
		timeUnix + ":" + tokenSecret
	return Token(base64.StdEncoding.EncodeToString([]byte(raw)))
}

func (token Token) decode() ([]string, bool) {
	decoded, err := base64.StdEncoding.DecodeString(string(token))
	if err != nil {
		return nil, false
	}

	parts := strings.Split(string(decoded), ":")
	if len(parts) != 5 {
		return nil, false
	}

	if parts[4] != tokenSecret {
		return nil, false
	}

	return parts, true
}

func (token Token) IsValid() bool {
	parts, ok := token.decode()
	if !ok {
		return false
	}

	if TokenStatus(parts[2]) != TokenValid {
		return false
	}

	return !token.IsExpired()
}

func (token Token) IsRevoked() bool {
	parts, ok := token.decode()
	if !ok {
		return false
	}
	return TokenStatus(parts[2]) == TokenRevoked
}

func (token Token) IsExpired() bool {
	parts, ok := token.decode()
	if !ok {
		return true
	}

	expUnix, err := strconv.ParseInt(parts[3], 10, 64)
	if err != nil {
		return true
	}

	return time.Now().Unix() > expUnix
}

func (token Token) UserId() string {
	parts, ok := token.decode()
	if !ok {
		return ""
	}
	return parts[0]
}

func (token Token) Email() string {
	parts, ok := token.decode()
	if !ok {
		return ""
	}
	return parts[1]
}

func (token Token) Status() string {
	parts, ok := token.decode()
	if !ok {
		return ""
	}
	return parts[2]
}
