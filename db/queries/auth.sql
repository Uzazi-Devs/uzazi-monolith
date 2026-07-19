-- name: GetUserByID :one
SELECT * FROM "user" WHERE "id" = $1;

-- name: GetSessionByToken :one
SELECT * FROM "session" WHERE "token" = $1;
