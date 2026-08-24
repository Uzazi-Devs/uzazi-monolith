package waitlist

import (
	"crypto/subtle"
	"net/http"
	"os"

	"github.com/Uzazi-Devs/uzazi-monolith/backend/internal/httpx"
)

// RequireAdmin gates a handler behind HTTP Basic Auth, checked against the
// ADMIN_USER/ADMIN_PASS env vars. There's no JWT/session verification
// anywhere in this backend yet, so Basic Auth is the smallest thing that
// keeps signup PII off the open internet until that exists.
func RequireAdmin(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		wantUser := os.Getenv("ADMIN_USER")
		wantPass := os.Getenv("ADMIN_PASS")

		user, pass, ok := r.BasicAuth()
		userMatch := subtle.ConstantTimeCompare([]byte(user), []byte(wantUser)) == 1
		passMatch := subtle.ConstantTimeCompare([]byte(pass), []byte(wantPass)) == 1

		if !ok || wantUser == "" || wantPass == "" || !userMatch || !passMatch {
			w.Header().Set("WWW-Authenticate", `Basic realm="waitlist admin"`)
			httpx.WriteError(w, http.StatusUnauthorized, "unauthorized")
			return
		}

		next(w, r)
	}
}
