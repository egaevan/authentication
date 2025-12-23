package usecase

import (
	"authentication/domain"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type Login struct {
	users map[string]*domain.User
}

func NewLoginUser(users map[string]*domain.User) *Login {
	return &Login{users: users}
}

var jwtSecret = []byte("secret")

func (uc Login) Execute(email string, password string) (token string, err error) {
	user, ok := uc.users[email]
	if !ok {
		return "", errors.New("invalid credentials")
	}

	if !user.IsActive() {
		return "", errors.New("user inactive")
	}

	passwordVO, err := domain.NewPassword(password)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	if passwordVO != user.Password() {
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
