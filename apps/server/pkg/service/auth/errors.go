package auth

import "errors"

var (
	GrantsNotFoundErr = errors.New("grants not found")
	InvalidGrantsErr  = errors.New("invalid grants")
)
