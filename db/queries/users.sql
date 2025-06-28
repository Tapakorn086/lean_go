-- name: GetAuthor :one
SELECT * FROM users WHERE id = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users ORDER BY id ASC;

-- name: CreateAuthor :one
INSERT INTO users (name, email, created_at)
VALUES ($1, $2, $3)
RETURNING *;

-- name: UpdateAuthor :exec
UPDATE users
SET email = $2, created_at = $3
WHERE id = $1;

-- name: DeleteAuthor :exec
DELETE FROM users WHERE id = $1;
