package auth

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/MicahParks/keyfunc/v3"
	"github.com/golang-jwt/jwt/v5"

	"github.com/deluxesande/uzazi-monolith/backend/internal/db"
)

// NewVerifier fetches the JWKS from the auth-service (e.g.
// http://auth-service:3000/api/auth/jwks). Returns an error if unreachable so
// the caller can decide whether that is fatal.
func NewVerifier(_ context.Context, jwksURL string, q *db.Queries) (*Verifier, error) {
	k, err := keyfunc.NewDefault([]string{jwksURL})
	if err != nil {
		return nil, err
	}
	return &Verifier{jwks: k, queries: q}, nil
}

// Verify returns the token subject (user id) if the token is valid.
func (v *Verifier) Verify(tokenStr string) (string, error) {
	tok, err := jwt.Parse(tokenStr, v.jwks.Keyfunc)
	if err != nil {
		return "", err
	}
	if !tok.Valid {
		return "", errors.New("invalid token")
	}
	sub, err := tok.Claims.GetSubject()
	if err != nil || sub == "" {
		return "", errors.New("missing subject")
	}
	return sub, nil
}

// Middleware requires a valid `Authorization: Bearer <jwt>` and stashes the
// resolved user in the request context.
func (v *Verifier) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		token := strings.TrimPrefix(header, "Bearer ")
		if token == "" || token == header {
			http.Error(w, "missing bearer token", http.StatusUnauthorized)
			return
		}
		userID, err := v.Verify(token)
		if err != nil {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}
		user, err := v.queries.GetUserByID(r.Context(), userID)
		if err != nil {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), ctxKey{}, user)))
	})
}

// UserFromContext returns the user set by Middleware.
func UserFromContext(ctx context.Context) (db.User, bool) {
	u, ok := ctx.Value(ctxKey{}).(db.User)
	return u, ok
}
