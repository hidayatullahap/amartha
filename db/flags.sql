-- name: InsertFlag :exec
INSERT INTO flags (key, value)
VALUES (?, ?);

-- name: GetFlag :one
SELECT key, value FROM flags
WHERE key = ?;

