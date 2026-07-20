package community

import (
	"context"

	"github.com/deluxesande/uzazi-monolith/backend/internal/db"
)

type Service interface {
	Post(ctx context.Context, userID, room, body string) (db.CommunityMessage, error)
	List(ctx context.Context, room string, limit int32) ([]db.CommunityMessage, error)
}

type service struct{ q *db.Queries }
