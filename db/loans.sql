-- name: GetLoan :one
SELECT * FROM loans
WHERE id = ? LIMIT 1;

-- name: ListLoans :many
SELECT * FROM loans;

-- name: CreateLoan :exec
INSERT INTO loans (id, borrower_id, principal_amount, rate, roi, agreement_letter_url)
VALUES (?, ?, ?, ?, ?, ?);

-- name: CreateLoanDetail :exec
INSERT INTO loan_details (loan_id, field_validator_id, visit_proof_url, approved_at, field_officer_id)
VALUES (?, ?, ?, ?, ?);

-- name: UpdateLoanPrincipleAmount :exec
UPDATE loans
SET principal_amount = principal_amount - ?,
    state = ?
WHERE id = ?;
