-- name: CreateWaitlistSignup :one
INSERT INTO waitlist_signup (name, email, stage, location, support)
VALUES ($1, $2, $3, $4, $5)
ON CONFLICT (lower(email)) DO NOTHING
RETURNING *;

-- name: ListWaitlistSignups :many
SELECT * FROM waitlist_signup
ORDER BY created_at DESC;

-- name: AcceptWaitlistSignup :exec
UPDATE waitlist_signup
SET status = 'accepted'
WHERE id = $1;
