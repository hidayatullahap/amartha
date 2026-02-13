package service

import "fmt"

type LoanService interface {
	GetLoans()
	GetLoan()
	CreateLoan()
	ApproveLoan()
	InvestLoan()
	DisburseLoan()
}

type loanService struct {
}

func NewLoanService() LoanService {
	return loanService{}
}

func (s loanService) ApproveLoan() {
	fmt.Println("approve loan")
}

func (s loanService) CreateLoan() {
	fmt.Println("create loan")
}

func (s loanService) DisburseLoan() {
	fmt.Println("disburse loan")
}

func (s loanService) InvestLoan() {
	fmt.Println("invest loan")
}

func (s loanService) GetLoans() {
	fmt.Println("get loans")
}

func (s loanService) GetLoan() {
	fmt.Println("get loan")
}
