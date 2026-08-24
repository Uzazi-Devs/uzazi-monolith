package waitlist

import (
	"net/http"

	"github.com/Uzazi-Devs/uzazi-monolith/backend/internal/httpx"
)

// List handles GET /admin/waitlist. Gated by RequireAdmin.
func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
	signups, err := h.Queries.ListWaitlistSignups(r.Context())
	if err != nil {
		httpx.WriteError(w, http.StatusInternalServerError, "could not load signups")
		return
	}

	httpx.WriteJSON(w, http.StatusOK, signups)
}
