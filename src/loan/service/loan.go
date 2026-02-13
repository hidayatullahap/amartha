package service

import (
	"amartha/generated/sqlc"
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
	queries *sqlc.Queries
	guard   GuardLoanService
}

func NewLoanService(queries *sqlc.Queries, guard GuardLoanService) LoanService {
	return loanService{queries, guard}
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
