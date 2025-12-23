package testdata

import "authentication/domain"

var (
	email, _    = domain.NewEmail("testing@gmail.com")
	email6, _   = domain.NewEmail("testing@gmail.com")
	password, _ = domain.NewPassword("Testing123*")

	Users = map[string]*domain.User{
		"testing@gmail.com":  domain.NewUser(email, password, "active"),
		"testing6@gmail.com": domain.NewUser(email6, password, "inactive"),
	}
)
