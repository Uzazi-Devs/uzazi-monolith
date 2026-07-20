package community

import (
	"context"

	"github.com/deluxesande/uzazi-monolith/backend/internal/db"
)

func NewService(q *db.Queries) Service { return &service{q: q} }

func (s *service) Post(ctx context.Context, userID, room, body string) (db.CommunityMessage, error) {
	return s.q.PostMessage(ctx, db.PostMessageParams{UserID: userID, Room: room, Body: body})
}

func (s *service) List(ctx context.Context, room string, limit int32) ([]db.CommunityMessage, error) {
	if limit <= 0 {
		limit = 50
	}
	return s.q.ListMessages(ctx, db.ListMessagesParams{Room: room, Limit: limit})
}
