package authentication

type User struct {
	email    string
	password string
	status   string
}

func (u User) Status() string {
	return u.status
}

func Register(email, password string) (*User, error) {
	return &User{
		email:    email,
		password: "bukanpassword",
		status:   "active",
	}, nil
}
