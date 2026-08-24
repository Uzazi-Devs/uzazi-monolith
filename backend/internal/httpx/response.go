// Package httpx holds small JSON response helpers shared by every handler
// package in this backend. Keep it dependency-free so any domain package
// can import it without pulling in unrelated packages.
package httpx

import (
	"encoding/json"
	"net/http"
)

func WriteJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}

func WriteError(w http.ResponseWriter, status int, msg string) {
	WriteJSON(w, status, map[string]string{"error": msg})
}

// WriteStatus writes {"status": status}, the shape most write endpoints
// return on success.
func WriteStatus(w http.ResponseWriter, code int, status string) {
	WriteJSON(w, code, map[string]string{"status": status})
}

// WriteIDStatus writes {"id": id, "status": status}, for endpoints that
// act on a path-scoped resource.
func WriteIDStatus(w http.ResponseWriter, code int, id, status string) {
	WriteJSON(w, code, map[string]string{"id": id, "status": status})
}
