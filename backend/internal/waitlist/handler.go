// Package waitlist handles marketing-site waitlist signups: a public submit
// endpoint and an internal admin list for reviewing and accepting them.
// Accepting a signup is informational only — it does not create or gate a
// user account; the operator emails the accepted person separately.
//
// JSON only, throughout — this is an API, not a web app.
package waitlist

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"

	"github.com/Uzazi-Devs/uzazi-monolith/backend/internal/db"
)

type Handler struct {
	Queries *db.Queries
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}

func writeError(w http.ResponseWriter, status int, msg string) {
	writeJSON(w, status, map[string]string{"error": msg})
}

type createSignupRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Stage    string `json:"stage"`
	Location string `json:"location"`
	Support  string `json:"support"`
	Consent  bool   `json:"consent"`
	Company  string `json:"company"` // honeypot — real users never fill this
}

// Create handles the public waitlist submit. Callers must POST JSON.
func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var req createSignupRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON body")
		return
	}

	// Honeypot: bots fill every field. Pretend success without writing.
	if req.Company != "" {
		writeJSON(w, http.StatusCreated, map[string]string{"status": "pending"})
		return
	}

	if req.Name == "" || req.Email == "" || req.Stage == "" || !req.Consent {
		writeError(w, http.StatusBadRequest, "missing required field")
		return
	}

	signup, err := h.Queries.CreateWaitlistSignup(r.Context(), db.CreateWaitlistSignupParams{
		Name:     req.Name,
		Email:    req.Email,
		Stage:    req.Stage,
		Location: req.Location,
		Support:  req.Support,
	})
	// ON CONFLICT DO NOTHING means an existing email returns no row —
	// pgx surfaces that as ErrNoRows. Treat a repeat signup as success too.
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		writeError(w, http.StatusInternalServerError, "could not save signup")
		return
	}

	writeJSON(w, http.StatusCreated, map[string]string{"status": signup.Status})
}

func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
	signups, err := h.Queries.ListWaitlistSignups(r.Context())
	if err != nil {
		writeError(w, http.StatusInternalServerError, "could not load signups")
		return
	}

	writeJSON(w, http.StatusOK, signups)
}

func (h *Handler) Accept(w http.ResponseWriter, r *http.Request) {
	var id pgtype.UUID
	if err := id.Scan(r.PathValue("id")); err != nil {
		writeError(w, http.StatusBadRequest, "bad id")
		return
	}

	if err := h.Queries.AcceptWaitlistSignup(r.Context(), id); err != nil {
		writeError(w, http.StatusInternalServerError, "could not accept signup")
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{"id": r.PathValue("id"), "status": "accepted"})
}
