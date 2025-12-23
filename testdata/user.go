package testdata

import (
	"authentication/domain"
	"strconv"
	"time"
)

var (
	email, _    = domain.NewEmail("testing@gmail.com")
	email6, _   = domain.NewEmail("testing@gmail.com")
	password, _ = domain.NewPassword("Testing123*")

	UserActive   = domain.NewUser(email, password, "active")
	UserInactive = domain.NewUser(email, password, "inactive")

	now = time.Now()

	Users = map[string]*domain.User{
		"testing@gmail.com":  UserActive,
		"testing6@gmail.com": UserInactive,
	}

	TokenValid   = domain.NewToken(UserActive.IdString(), UserActive.Email(), domain.TokenValid, "")
	TokenExpired = domain.NewToken(UserActive.IdString(), UserActive.Email(), domain.TokenExpired, strconv.FormatInt(now.AddDate(0, 0, -1).Unix(), 10))
	TokenRevoked = domain.NewToken(UserActive.IdString(), UserActive.Email(), domain.TokenRevoked, "")
)
