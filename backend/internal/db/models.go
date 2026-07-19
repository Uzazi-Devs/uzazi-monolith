package db

import "github.com/jackc/pgx/v5/pgtype"

// ponytail: these structs are hand-written to match `sqlc generate` output so
// the backend compiles before sqlc is run. Regenerate from db/ with
// `sqlc generate`; CI fails if the committed code drifts from the migrations.

type User struct {
	ID            string             `json:"id"`
	Name          string             `json:"name"`
	Email         string             `json:"email"`
	EmailVerified bool               `json:"emailVerified"`
	Image         pgtype.Text        `json:"image"`
	CreatedAt     pgtype.Timestamptz `json:"createdAt"`
	UpdatedAt     pgtype.Timestamptz `json:"updatedAt"`
}

type Session struct {
	ID        string             `json:"id"`
	UserID    string             `json:"userId"`
	Token     string             `json:"token"`
	ExpiresAt pgtype.Timestamptz `json:"expiresAt"`
	IpAddress pgtype.Text        `json:"ipAddress"`
	UserAgent pgtype.Text        `json:"userAgent"`
	CreatedAt pgtype.Timestamptz `json:"createdAt"`
	UpdatedAt pgtype.Timestamptz `json:"updatedAt"`
}

type HealthRecord struct {
	ID          pgtype.UUID        `json:"id"`
	UserID      string             `json:"user_id"`
	Kind        string             `json:"kind"`
	TriageLevel pgtype.Text        `json:"triage_level"`
	Notes       string             `json:"notes"`
	CreatedAt   pgtype.Timestamptz `json:"created_at"`
}

type CommunityMessage struct {
	ID        pgtype.UUID        `json:"id"`
	UserID    string             `json:"user_id"`
	Room      string             `json:"room"`
	Body      string             `json:"body"`
	CreatedAt pgtype.Timestamptz `json:"created_at"`
}

type MediaObject struct {
	ID          pgtype.UUID        `json:"id"`
	UserID      string             `json:"user_id"`
	Bucket      string             `json:"bucket"`
	ObjectKey   string             `json:"object_key"`
	ContentType string             `json:"content_type"`
	CreatedAt   pgtype.Timestamptz `json:"created_at"`
}
