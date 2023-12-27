package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AccessToken struct {
	grants *Grants
	rawJWT string
	issuer string
	secret string
}

func (t *AccessToken) ToJWT() (string, error) {
	// send created
	if t.rawJWT != "" {
		return t.rawJWT, nil
	}

	claims := &JWTClaims{
		Grants: t.grants,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    t.issuer,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TokenExpiration)),
		},
	}

	// create a new one
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	rawJWT, err := token.SignedString([]byte(t.secret))

	if err != nil {
		return "", err
	}

	return rawJWT, nil
}
