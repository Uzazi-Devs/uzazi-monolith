package health

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"

	"github.com/deluxesande/uzazi/backend/internal/db"
)

// Service is how other modules use health records — a Go interface, no HTTP.
type Service interface {
	CreateRecord(ctx context.Context, userID, kind, triageLevel, notes string) (db.HealthRecord, error)
	ListForUser(ctx context.Context, userID string) ([]db.HealthRecord, error)
}

type service struct{ q *db.Queries }

func NewService(q *db.Queries) Service { return &service{q: q} }

func (s *service) CreateRecord(ctx context.Context, userID, kind, triageLevel, notes string) (db.HealthRecord, error) {
	tl := pgtype.Text{}
	if triageLevel != "" {
		tl = pgtype.Text{String: triageLevel, Valid: true}
	}
	return s.q.CreateHealthRecord(ctx, db.CreateHealthRecordParams{
		UserID: userID, Kind: kind, TriageLevel: tl, Notes: notes,
	})
}

func (s *service) ListForUser(ctx context.Context, userID string) ([]db.HealthRecord, error) {
	return s.q.ListHealthRecordsByUser(ctx, userID)
}
