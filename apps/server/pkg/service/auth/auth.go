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

type AccessToken struct {
	grants *Grants
	rawJWT string
	issuer string
	secret string
}

type Grants struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	RoomID    string `json:"roomId"`
	RoomAdmin bool   `json:"roomAdmin"`
}

type JWTClaims struct {
	Grants *Grants `json:"grants"`

	jwt.RegisteredClaims
}

func (s *AuthService) NewAccessTokenFromGrants(grants *Grants) *AccessToken {
	return &AccessToken{
		grants: grants,
		issuer: s.conf.Issuer,
		secret: s.conf.Secret,
	}
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
