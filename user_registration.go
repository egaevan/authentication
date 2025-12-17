package authentication

import "golang.org/x/crypto/bcrypt"

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
