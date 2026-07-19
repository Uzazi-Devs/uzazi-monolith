-- name: CreateMediaObject :one
INSERT INTO media_object (user_id, bucket, object_key, content_type)
VALUES ($1, $2, $3, $4)
RETURNING *;
