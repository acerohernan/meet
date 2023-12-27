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
		// room service open endpoints
		"/twirp/twirp.v1.RoomService/CreateRoom": true,
	}
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
		handleError(w, "acess token not found")
		return
	}

	at, err := m.svc.NewAccessTokenFromRawJWT(authToken)

	if err != nil {
		handleError(w, "could not parse access token")
		return
	}

	// set grants in context
	ctx := r.Context()
	r = r.WithContext(context.WithValue(ctx, "grants", at.grants))

	next(w, r)
}

func handleError(w http.ResponseWriter, msg string) {
	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte(msg))
}
