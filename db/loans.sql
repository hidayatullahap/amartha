-- name: GetLoan :one
SELECT * FROM loans
WHERE id = ? LIMIT 1;

-- name: ListLoans :many
SELECT * FROM loans;

-- name: CreateLoan :exec
INSERT INTO loans (id, borrower_id, principal_amount, rate, roi, agreement_letter_url)
VALUES (?, ?, ?, ?, ?, ?);
