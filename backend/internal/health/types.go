package health

import (
	"context"

	"github.com/Uzazi-Devs/uzazi-monolith/backend/internal/db"
)

// Service is how other modules use health records — a Go interface, no HTTP.
type Service interface {
	CreateRecord(ctx context.Context, userID, kind, triageLevel, notes string) (db.HealthRecord, error)
	ListForUser(ctx context.Context, userID string) ([]db.HealthRecord, error)
}

type service struct{ q *db.Queries }
