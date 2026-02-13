package service

import (
	"amartha/src/loan/request"
	"context"
	"fmt"
	"log"
)

type LoanService interface {
	GetLoans()
	GetLoan()
	CreateLoan(ctx context.Context, req request.CreateLoanRequest) (*request.CreateLoanRequest, error)
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

func (s loanService) CreateLoan(ctx context.Context, req request.CreateLoanRequest) (*request.CreateLoanRequest, error) {
	log.Println(req)
	return &req, nil
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
