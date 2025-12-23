package authentication

import "errors"

func Logout(token string) error {
	if token != "token" {
		return errors.New("invalid token")
	}

	return nil
}
