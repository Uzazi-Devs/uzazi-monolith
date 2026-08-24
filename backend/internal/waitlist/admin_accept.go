package waitlist

import (
	"net/http"

	"github.com/jackc/pgx/v5/pgtype"

	"github.com/Uzazi-Devs/uzazi-monolith/backend/internal/httpx"
)

// Accept handles POST /admin/waitlist/{id}/accept. Gated by RequireAdmin.
func (h *Handler) Accept(w http.ResponseWriter, r *http.Request) {
	var id pgtype.UUID
	if err := id.Scan(r.PathValue("id")); err != nil {
		httpx.WriteError(w, http.StatusBadRequest, "bad id")
		return
	}

	if err := h.Queries.AcceptWaitlistSignup(r.Context(), id); err != nil {
		httpx.WriteError(w, http.StatusInternalServerError, "could not accept signup")
		return
	}

	httpx.WriteIDStatus(w, http.StatusOK, r.PathValue("id"), "accepted")
}
