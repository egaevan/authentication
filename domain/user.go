package domain

import (
	"math/rand"
	"strconv"
)

type User struct {
	id       int
	email    Email
	password Password
	status   string
}

func NewUser(email Email, passwordHash Password, status string) *User {
	if status == "" {
		status = "active"
	}
	return &User{
		id:       rand.Int(),
		email:    email,
		password: passwordHash,
		status:   status,
	}
}

func (u User) Status() string {
	return u.status
}

func (u User) IsActive() bool {
	return u.status == "active"
}

func (u User) Email() Email {
	return u.email
}

func (u User) Password() Password {
	return u.password
}

func (u User) ID() int {
	return u.id
}

func (u User) IdString() string {
	return strconv.Itoa(u.id)
}
