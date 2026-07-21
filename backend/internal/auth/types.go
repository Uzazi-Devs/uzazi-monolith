package auth

import (
	"github.com/MicahParks/keyfunc/v3"

	"github.com/Uzazi-Devs/uzazi-monolith/backend/internal/db"
)

// Verifier validates BetterAuth-issued JWTs. BetterAuth (services/auth-service)
// is the ONLY credential issuer in the stack. This package never signs or
// issues anything — it checks signature + expiry against the auth-service JWKS
// and loads the user through sqlc-generated queries.
type Verifier struct {
	jwks    keyfunc.Keyfunc
	queries *db.Queries
}

// ctxKey is the private context key under which Middleware stores the
// resolved user.
type ctxKey struct{}
