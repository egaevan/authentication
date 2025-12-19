package domain

import (
	"errors"
	"regexp"
)

type Email string

func NewEmail(value string) (Email, error) {
	var regex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	if !regex.MatchString(value) {
		return "", errors.New("invalid email format")
	}
	return Email(value), nil
}
