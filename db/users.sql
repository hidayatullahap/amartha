-- name: GetUserByUsername :one
SELECT id, username, role FROM users
WHERE username = ?;

-- name: GetUserById :one
SELECT id, username, role FROM users
WHERE id = ?;