package domain

type TokenStatus string

const (
	TokenValid   TokenStatus = "valid"
	TokenExpired TokenStatus = "expired"
	TokenRevoked TokenStatus = "revoked"
)
