package response

type CreateLoanResponse struct {
	ID string `json:"id"`
}

type CreateLoanDetailResponse struct {
	LoanID string `json:"loan_id"`
}

type CreateLoanInvestResponse struct {
	LoanID       string `json:"loan_id"`
	InvestmentId string `json:"investment_id"`
}
