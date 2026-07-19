-- name: CreateHealthRecord :one
INSERT INTO health_record (user_id, kind, triage_level, notes)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: ListHealthRecordsByUser :many
SELECT * FROM health_record
WHERE user_id = $1
ORDER BY created_at DESC;
