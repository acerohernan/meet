package auth

import (
	"time"

	"github.com/acerohernan/meet/pkg/config"
	"github.com/golang-jwt/jwt/v5"
)

var (
	TokenExpiration = time.Minute * 5
)

type AuthService struct {
	conf *config.JWTConfig
}

func NewAuthService(conf *config.Config) *AuthService {
	return &AuthService{
		conf: conf.JWT,
	}
}

func (s *AuthService) NewAccessTokenFromGrants(grants *Grants) *AccessToken {
	return &AccessToken{
		grants: grants,
		issuer: s.conf.Issuer,
		secret: s.conf.Secret,
	}
}

type JWTClaims struct {
	Grants *Grants `json:"grants"`

	jwt.RegisteredClaims
}

func (s *AuthService) NewAccessTokenFromRawJWT(rawJWT string) (*AccessToken, error) {
	token, err := jwt.ParseWithClaims(rawJWT, &JWTClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(s.conf.Secret), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*JWTClaims)

	if !ok {
		return nil, err
	}

	return &AccessToken{
		grants: claims.Grants,
		rawJWT: rawJWT,
		issuer: s.conf.Issuer,
		secret: s.conf.Secret,
	}, nil
}
