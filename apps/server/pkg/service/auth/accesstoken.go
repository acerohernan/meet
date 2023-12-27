package auth

type AccessToken struct {
	grants *Grants
	rawJWT string
	issuer string
	secret string
}

func (t *AccessToken) ToJWT() string {
	return t.rawJWT
}
