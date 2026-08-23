package waitlist

import (
	"net/http"
	"os"
	"strings"
)

// WithCORS allows a small set of origins (comma-separated in the
// CORS_ALLOWED_ORIGINS env var) to call a JSON endpoint via fetch/XHR.
// Only the public /waitlist submit needs this — it's the only endpoint
// called from browser JS on another origin (the marketing site).
func WithCORS(next http.HandlerFunc) http.HandlerFunc {
	allowed := strings.Split(os.Getenv("CORS_ALLOWED_ORIGINS"), ",")

	return func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		for _, a := range allowed {
			if a != "" && a == origin {
				w.Header().Set("Access-Control-Allow-Origin", origin)
				w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
				w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
				break
			}
		}

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next(w, r)
	}
}
