-- name: GetLoan :one
SELECT * FROM loans
WHERE id = ? LIMIT 1;

-- name: ListLoans :many
SELECT * FROM loans;
