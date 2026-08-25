package waitlist

import (
	"net/http"
	"strings"

	"github.com/Uzazi-Devs/uzazi-monolith/backend/internal/authjwt"
	"github.com/Uzazi-Devs/uzazi-monolith/backend/internal/httpx"
)

// RequireAdmin gates a handler behind a BetterAuth JWT (Authorization:
// Bearer <token>, verified against AUTH_JWKS_URL) whose role claim is
// "admin" — the role services/auth-service's admin() plugin assigns.
func RequireAdmin(verifier *authjwt.Verifier) func(http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			token, ok := strings.CutPrefix(r.Header.Get("Authorization"), "Bearer ")
			if !ok || token == "" {
				w.Header().Set("WWW-Authenticate", `Bearer realm="waitlist admin"`)
				httpx.WriteError(w, http.StatusUnauthorized, "unauthorized")
				return
			}

			claims, err := verifier.Verify(token)
			if err != nil || claims.Role != "admin" {
				w.Header().Set("WWW-Authenticate", `Bearer realm="waitlist admin"`)
				httpx.WriteError(w, http.StatusUnauthorized, "unauthorized")
				return
			}

			next(w, r)
		}
	}
}
