package auth

import (
	"time"

	"github.com/acerohernan/meet/pkg/config"
	"github.com/golang-jwt/jwt/v5"
)

var (
	TokenExpiration = time.Hour * 5
)

type AuthService struct {
	conf *config.JWTConfig
}

func NewAuthService(conf *config.Config) *AuthService {
	return &AuthService{
		conf: conf.JWT,
	}
}

type JWTClaims struct {
	Grants *Grants `json:"grants"`

	jwt.RegisteredClaims
}

func (s *AuthService) NewAccessTokenFromGrants(grants *Grants) (*AccessToken, error) {
	claims := &JWTClaims{
		Grants: grants,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    s.conf.Issuer,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TokenExpiration)),
		},
	}

	// create a new one
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	rawJWT, err := token.SignedString([]byte(s.conf.Secret))

	if err != nil {
		return nil, err
	}

	return &AccessToken{
		grants: grants,
		issuer: s.conf.Issuer,
		secret: s.conf.Secret,
		rawJWT: rawJWT,
	}, nil
}

func (s *AuthService) NewAccessTokenFromRawJWT(rawJWT string) (*AccessToken, error) {
	token, err := jwt.ParseWithClaims(rawJWT, &JWTClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(s.conf.Secret), nil
	}, jwt.WithExpirationRequired())

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
