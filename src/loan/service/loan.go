package service

import (
	"amartha/generated/sqlc"
	"amartha/src/loan/constants"
	"amartha/src/loan/request"
	"amartha/src/loan/response"
	uerror "amartha/src/utils/error"
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type LoanService interface {
	GetLoans()
	GetLoan()
	CreateLoan(ctx context.Context, req request.CreateLoanRequest) (*response.CreateLoanResponse, error)
	ApproveLoan(ctx context.Context, req request.CreateLoanDetailRequest) (*response.CreateLoanDetailResponse, error)
	InvestLoan(ctx context.Context, req request.CreateInvestRequest) (*response.CreateLoanInvestResponse, error)
	DisburseLoan(ctx context.Context, req request.DisburseLoanRequest) error
}

type loanService struct {
	db      *sql.DB
	queries *sqlc.Queries
}

func NewLoanService(db *sql.DB, queries *sqlc.Queries) LoanService {
	return loanService{db, queries}
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
	})
	if err != nil {
		return nil, err
	}

	return &response.CreateLoanDetailResponse{
		LoanID: req.LoanID,
	}, nil
}

func (s loanService) InvestLoan(ctx context.Context, req request.CreateInvestRequest) (*response.CreateLoanInvestResponse, error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	qtx := s.queries.WithTx(tx)

	loan, err := qtx.GetLoan(ctx, req.LoanID)
	if err != nil {
		return nil, err
	}

	currentTotal, err := qtx.GetTotalInvestment(ctx, req.LoanID)
	if err != nil {
		return nil, err
	}

	newTotal := currentTotal.Float64 + float64(req.Amount)
	if newTotal > loan.PrincipalAmount {
		return nil, uerror.ErrExceedsPrincipleAmount
	}

	investmentID := uuid.NewString()
	err = qtx.CreateInvestment(ctx, sqlc.CreateInvestmentParams{
		ID:         investmentID,
		LoanID:     req.LoanID,
		InvestorID: req.InvestorId,
		Amount:     float64(req.Amount),
	})
	if err != nil {
		return nil, err
	}

	if newTotal == loan.PrincipalAmount {
		err = qtx.UpdateLoanState(ctx, sqlc.UpdateLoanStateParams{
			ID:    req.LoanID,
			State: sql.NullString{Valid: true, String: constants.StateInvested.String()},
		})
		if err != nil {
			return nil, err
		}

		// TODO: Trigger email send
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return &response.CreateLoanInvestResponse{
		LoanID:       req.LoanID,
		InvestmentId: investmentID,
	}, nil
}

func (s loanService) DisburseLoan(ctx context.Context, req request.DisburseLoanRequest) error {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	qtx := s.queries.WithTx(tx)

	err = qtx.UpdateLoanState(ctx, sqlc.UpdateLoanStateParams{
		ID:    req.LoanID,
		State: sql.NullString{Valid: true, String: constants.StateDisbursed.String()},
	})
	if err != nil {
		return err
	}

	err = qtx.UpdateDisbursementDetails(ctx, sqlc.UpdateDisbursementDetailsParams{
		LoanID:             req.LoanID,
		FieldOfficerID:     sql.NullInt64{Int64: req.FieldOfficerID, Valid: true},
		SignedAgreementUrl: sql.NullString{String: req.SignedAgreementURL, Valid: true},
		DisbursedAt:        sql.NullTime{Time: time.Now(), Valid: true},
	})
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (s loanService) GetLoans() {
	fmt.Println("get loans")
}

func (s loanService) GetLoan() {
	fmt.Println("get loan")
}
