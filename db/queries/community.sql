-- name: PostMessage :one
INSERT INTO community_message (user_id, room, body)
VALUES ($1, $2, $3)
RETURNING *;

-- name: ListMessages :many
SELECT * FROM community_message
WHERE room = $1
ORDER BY created_at DESC
LIMIT $2;
