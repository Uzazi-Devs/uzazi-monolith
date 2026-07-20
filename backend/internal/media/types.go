package media

import (
	"context"

	"github.com/Uzazi-Devs/uzazi-monolith/backend/internal/db"
)

// Service records Cloud Storage object references. Actual byte uploads go
// straight to the bucket (signed URL); we only persist the ref.
type Service interface {
	Register(ctx context.Context, userID, bucket, objectKey, contentType string) (db.MediaObject, error)
}

type service struct{ q *db.Queries }
