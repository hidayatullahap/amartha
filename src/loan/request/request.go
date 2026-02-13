package request

import "time"

type CreateLoanRequest struct {
	BorrowerId         int64   `json:"borrower_id"`
	PrincipalAmount    float64 `json:"principal_amount"`
	Rate               float64 `json:"rate"`
	Roi                float64 `json:"roi"`
	AgreementLetterUrl string  `json:"agreement_letter_url"`
}

type CreateLoanDetailRequest struct {
	LoanID           string    `json:"-"`
	FieldValidatorID int64     `json:"field_validator_id"`
	VisitProofUrl    string    `json:"visit_proof_url"`
	FieldOfficerID   int64     `json:"field_officer_id"`
	ApprovalDate     time.Time `json:"approval_date"`
}

type CreateInvestRequest struct {
	LoanID     string `json:"-"`
	Amount     int64  `json:"amount"`
	InvestorId int64  `json:"-"`
}
