package service

import (
	"amartha/generated/sqlc"
	"amartha/src/loan/request"
	"amartha/src/loan/response"
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

type LoanService interface {
	GetLoans()
	GetLoan()
	CreateLoan(ctx context.Context, req request.CreateLoanRequest) (*response.CreateLoanResponse, error)
	ApproveLoan()
	InvestLoan()
	DisburseLoan()
}

type loanService struct {
	queries *sqlc.Queries
}

func NewLoanService(queries *sqlc.Queries) LoanService {
	return loanService{queries}
}

func (s loanService) ApproveLoan() {
	fmt.Println("approve loan")
}

func (s loanService) CreateLoan(ctx context.Context, req request.CreateLoanRequest) (*response.CreateLoanResponse, error) {
	id := uuid.NewString()
	err := s.queries.CreateLoan(ctx, sqlc.CreateLoanParams{
		ID:                 id,
		BorrowerID:         req.BorrowerId,
		PrincipalAmount:    req.PrincipalAmount,
		Rate:               req.Rate,
		Roi:                req.Roi,
		AgreementLetterUrl: sql.NullString{Valid: true, String: req.AgreementLetterUrl},
	})
	if err != nil {
		return nil, err
	}

	return &response.CreateLoanResponse{
		ID: id,
	}, nil
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
