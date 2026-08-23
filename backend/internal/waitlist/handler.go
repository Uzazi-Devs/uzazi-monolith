// Package waitlist handles marketing-site waitlist signups: a public submit
// endpoint and an internal admin list for reviewing and accepting them.
// Accepting a signup is informational only — it does not create or gate a
// user account; the operator emails the accepted person separately.
package waitlist

import (
	"errors"
	"html/template"
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"

	"github.com/Uzazi-Devs/uzazi-monolith/backend/internal/db"
)

type Handler struct {
	Queries *db.Queries
}

// Create handles the public waitlist form POST. It's a plain HTML form
// submission (browser navigation), not a fetch/XHR, so no CORS handling is
// needed here.
func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "bad form", http.StatusBadRequest)
		return
	}

	// Honeypot: bots fill every field. Pretend success without writing.
	if r.FormValue("company") != "" {
		http.Redirect(w, r, "/waitlist/thanks", http.StatusSeeOther)
		return
	}

	name := r.FormValue("name")
	email := r.FormValue("email")
	stage := r.FormValue("stage")
	consent := r.FormValue("consent")
	if name == "" || email == "" || stage == "" || consent == "" {
		http.Error(w, "missing required field", http.StatusBadRequest)
		return
	}

	_, err := h.Queries.CreateWaitlistSignup(r.Context(), db.CreateWaitlistSignupParams{
		Name:     name,
		Email:    email,
		Stage:    stage,
		Location: r.FormValue("location"),
		Support:  r.FormValue("support"),
	})
	// ON CONFLICT DO NOTHING means an existing email returns no row —
	// pgx surfaces that as ErrNoRows. Treat a repeat signup as success too.
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		http.Error(w, "could not save signup", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/waitlist/thanks", http.StatusSeeOther)
}

var listTemplate = template.Must(template.New("list").Parse(`<!doctype html>
<html>
<head><title>Waitlist</title></head>
<body style="font-family: system-ui, sans-serif; margin: 2rem;">
<h1>Waitlist signups</h1>
<table border="1" cellpadding="8" style="border-collapse: collapse;">
<tr><th>Created</th><th>Name</th><th>Email</th><th>Stage</th><th>Location</th><th>Support</th><th>Status</th><th></th></tr>
{{range .}}
<tr>
<td>{{.CreatedAt.Time.Format "2006-01-02 15:04"}}</td>
<td>{{.Name}}</td>
<td>{{.Email}}</td>
<td>{{.Stage}}</td>
<td>{{.Location}}</td>
<td>{{.Support}}</td>
<td>{{.Status}}</td>
<td>
{{if eq .Status "pending"}}
<form method="post" action="/admin/waitlist/{{.ID.Value}}/accept">
<button type="submit">Accept</button>
</form>
{{end}}
</td>
</tr>
{{end}}
</table>
</body>
</html>
`))

func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
	signups, err := h.Queries.ListWaitlistSignups(r.Context())
	if err != nil {
		http.Error(w, "could not load signups", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := listTemplate.Execute(w, signups); err != nil {
		http.Error(w, "render error", http.StatusInternalServerError)
	}
}

func (h *Handler) Accept(w http.ResponseWriter, r *http.Request) {
	var id pgtype.UUID
	if err := id.Scan(r.PathValue("id")); err != nil {
		http.Error(w, "bad id", http.StatusBadRequest)
		return
	}

	if err := h.Queries.AcceptWaitlistSignup(r.Context(), id); err != nil {
		http.Error(w, "could not accept signup", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/admin/waitlist", http.StatusSeeOther)
}
