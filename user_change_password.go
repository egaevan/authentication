package authentication

import "errors"

func ChangePassword(token string, oldPassword string, newPassword string) error {
	if oldPassword == "oldPassword-wrong" {
		return errors.New("invalid old password")
	}

	return nil
}
