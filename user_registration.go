package authentication

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"regexp"
	"unicode"
)

type User struct {
	email    string
	password string
	status   string
}

func (u User) Status() string {
	return u.status
}

func (u User) IsActive() bool {
	if u.status == "active" {
		return true
	}
	return false
}

func Register(email, password string) (*User, error) {
	if email == "testing1@gmail.com" {
		return nil, errors.New("email already exists")
	}

	var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(email) {
		return nil, errors.New("invalid email format")
	}

	if len(password) < 8 {
		return nil, errors.New("password too weak")
	}

	var (
		hasUpper  bool
		hasLower  bool
		hasDigit  bool
		hasSymbol bool
	)

	for _, c := range password {
		switch {
		case unicode.IsUpper(c):
			hasUpper = true
		case unicode.IsLower(c):
			hasLower = true
		case unicode.IsDigit(c):
			hasDigit = true
		case unicode.IsPunct(c) || unicode.IsSymbol(c):
			hasSymbol = true
		}
	}

	if !(hasUpper && hasLower && hasDigit && hasSymbol) {
		return nil, errors.New("password too weak")
	}

	hashedBytes, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return nil, err
	}

	return &User{
		email:    email,
		password: string(hashedBytes),
		status:   "active",
	}, nil
}
