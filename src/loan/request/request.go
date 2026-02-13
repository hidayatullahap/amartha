package request

type CreateLoanRequest struct {
	BorrowerId         int64   `json:"borrower_id"`
	PrincipalAmount    float64 `json:"principal_amount"`
	Rate               float64 `json:"rate"`
	Roi                float64 `json:"roi"`
	AgreementLetterUrl string  `json:"agreement_letter_url"`
}
