-- name: GetUser :one
SELECT id, username, role FROM users
WHERE username = ?;