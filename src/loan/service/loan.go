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
	ApproveLoan(ctx context.Context, req request.CreateLoanDetailRequest) (*response.CreateLoanDetailResponse, error)
	InvestLoan()
	DisburseLoan()
}

type loanService struct {
	queries *sqlc.Queries
}

func NewLoanService(queries *sqlc.Queries) LoanService {
	return loanService{queries}
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

func (s loanService) ApproveLoan(ctx context.Context, req request.CreateLoanDetailRequest) (*response.CreateLoanDetailResponse, error) {
	err := s.queries.CreateLoanDetail(ctx, sqlc.CreateLoanDetailParams{
		LoanID:           req.LoanID,
		FieldValidatorID: sql.NullInt64{Valid: true, Int64: req.FieldValidatorID},
		VisitProofUrl:    sql.NullString{Valid: true, String: req.VisitProofUrl},
		ApprovedAt:       sql.NullTime{Valid: true, Time: req.ApprovalDate},
		FieldOfficerID:   sql.NullInt64{Valid: true, Int64: req.FieldOfficerID},
	})
	if err != nil {
		return nil, err
	}

	return &response.CreateLoanDetailResponse{
		LoanID: req.LoanID,
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
