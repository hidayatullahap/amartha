-- name: GetLoan :one
SELECT * FROM loans
WHERE id = ? LIMIT 1;

-- name: ListLoans :many
SELECT * FROM loans;

-- name: CreateLoan :exec
INSERT INTO loans (id, borrower_id, principal_amount, rate, roi, agreement_letter_url)
VALUES (?, ?, ?, ?, ?, ?);

-- name: CreateLoanDetail :exec
INSERT INTO loan_details (loan_id, field_validator_id, visit_proof_url, approved_at)
VALUES (?, ?, ?, ?);

-- name: UpdateLoanPrincipleAmount :exec
UPDATE loans
SET principal_amount = principal_amount - ?
WHERE id = ?;

-- name: CreateInvestment :exec
INSERT INTO investments (id, loan_id, investor_id, amount)
VALUES (?, ?, ?, ?);

-- name: GetTotalInvestment :one
SELECT SUM(amount) AS total_invested 
FROM investments 
WHERE loan_id = ?;

-- name: UpdateLoanState :exec
UPDATE loans
SET state = ?
WHERE id = ?;

-- name: UpdateDisbursementDetails :exec
UPDATE loan_details
SET 
    field_officer_id = ?,
    signed_agreement_url = ?,
    disbursed_at = ?
WHERE loan_id = ?;

-- name: GetInvestorEmailsByLoanID :many
SELECT DISTINCT u.email
FROM investments i
JOIN users u ON i.investor_id = u.id
WHERE i.loan_id = ?;
