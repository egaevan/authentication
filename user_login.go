package authentication

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var jwtSecret = []byte("secret")

func Login(email string, password string) (token string, err error) {
	if email != "testing@gmail.com" {
		return "", errors.New("invalid credentials")
	}

	if password != "Testing123*" {
		return "", errors.New("invalid credentials")
	}

	claims := jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(24 * time.Hour).Unix(),
		"iat":   time.Now().Unix(),
	}

	tokenJwt := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return tokenJwt.SignedString(jwtSecret)
}
