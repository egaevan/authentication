package authentication

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var jwtSecret = []byte("secret")

func Login(email string, password string) (token string, err error) {
	claims := jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(24 * time.Hour).Unix(),
		"iat":   time.Now().Unix(),
	}

	tokenJwt := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return tokenJwt.SignedString(jwtSecret)
}
