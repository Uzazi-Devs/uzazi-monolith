package media

import (
	"context"

	"github.com/deluxesande/uzazi/backend/internal/db"
)

// Service records Cloud Storage object references. Actual byte uploads go
// straight to the bucket (signed URL); we only persist the ref.
type Service interface {
	Register(ctx context.Context, userID, bucket, objectKey, contentType string) (db.MediaObject, error)
}

type service struct{ q *db.Queries }

func NewService(q *db.Queries) Service { return &service{q: q} }

func (s *service) Register(ctx context.Context, userID, bucket, objectKey, contentType string) (db.MediaObject, error) {
	return s.q.CreateMediaObject(ctx, db.CreateMediaObjectParams{
		UserID: userID, Bucket: bucket, ObjectKey: objectKey, ContentType: contentType,
	})
}
