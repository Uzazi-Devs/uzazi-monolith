package media

import (
	"context"

	"github.com/deluxesande/uzazi-monolith/backend/internal/db"
)

func NewService(q *db.Queries) Service { return &service{q: q} }

func (s *service) Register(ctx context.Context, userID, bucket, objectKey, contentType string) (db.MediaObject, error) {
	return s.q.CreateMediaObject(ctx, db.CreateMediaObjectParams{
		UserID: userID, Bucket: bucket, ObjectKey: objectKey, ContentType: contentType,
	})
}
