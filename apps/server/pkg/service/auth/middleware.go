package auth

import (
	"context"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	bearerPrefix        = "Bearer "
	accessTokenParam    = "access_token"
)

var (
	PublicRoutes = map[string]bool{
		"/": true,
		// guest join request endpoint
		"/join": true,
		// room service open endpoints
		"/twirp/twirp.v1.RoomService/CreateRoom": true,
		"/twirp/twirp.v1.RoomService/VerifyRoom": true,
	}
	GrantsCTXKey = "grants"
)

type AuthMiddleware struct {
	svc *AuthService
}

func NewAuthMiddleware(svc *AuthService) *AuthMiddleware {
	return &AuthMiddleware{
		svc: svc,
	}
}

func (m *AuthMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	if PublicRoutes[r.URL.Path] {
		next(w, r)
		return
	}

	// extract jwt
	authHeader := r.Header.Get(authorizationHeader)
	var authToken string

	if authHeader != "" {
		if !strings.HasPrefix(authHeader, bearerPrefix) {
			handleError(w, "invalid access token")
			return
		}

		authToken = authHeader[len(bearerPrefix):]
	} else {
		// attempt to find from request header
		authToken = r.FormValue(accessTokenParam)
	}

	if authToken == "" {
		handleError(w, "access token not found")
		return
	}

	at, err := m.svc.NewAccessTokenFromRawJWT(authToken)

	if err != nil {
		handleError(w, "could not parse access token")
		return
	}

	// set grants in context
	ctx := r.Context()
	r = r.WithContext(context.WithValue(ctx, GrantsCTXKey, at.grants))

	next(w, r)
}

func handleError(w http.ResponseWriter, msg string) {
	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte(msg))
}

func GetGrantsFromCTX(ctx context.Context) (*Grants, error) {
	ctxGrants := ctx.Value(GrantsCTXKey)
	if ctxGrants == nil {
		return nil, GrantsNotFoundErr
	}

	grants, ok := ctxGrants.(*Grants)

	if !ok {
		return nil, InvalidGrantsErr
	}

	return grants, nil
}
