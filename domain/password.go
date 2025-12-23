package domain

import (
	"errors"
	"unicode"
)

type Password string

func NewPassword(rawPass string) (Password, error) {
	var (
		hasUpper  bool
		hasLower  bool
		hasDigit  bool
		hasSymbol bool
	)

	if len(rawPass) < 8 {
		return "", errors.New("password too weak")
	}

	for _, c := range rawPass {
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
		return "", errors.New("password too weak")
	}

	//hashedBytes, err := bcrypt.GenerateFromPassword(
	//	[]byte(rawPass),
	//	bcrypt.DefaultCost,
	//)
	//if err != nil {
	//	return "", err
	//}

	return Password(rawPass), nil
}
