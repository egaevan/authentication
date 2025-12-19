package domain

type User struct {
	email    Email
	password Password
	status   string
}

func NewUser(email Email, passwordHash Password) (*User, error) {
	return &User{
		email:    email,
		password: passwordHash,
		status:   "active",
	}, nil
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

func (u User) Email() Email {
	return u.email
}

func (u User) Password() Password {
	return u.password
}
