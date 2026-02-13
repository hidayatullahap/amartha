package service

import (
	"amartha/generated/sqlc"
	"amartha/src/loan/request"
	"amartha/src/utils/event"
	"context"
	"errors"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestCreateLoan(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		mockQueries := &mockLoanQueries{}
		svc := &loanService{
			queries: mockQueries,
		}

		req := request.CreateLoanRequest{
			BorrowerId:      123,
			PrincipalAmount: 5000000,
			Rate:            10.0,
			Roi:             12.0,
		}

		resp, err := svc.CreateLoan(context.Background(), req)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if resp == nil || resp.ID == "" {
			t.Error("expected a valid response with a generated ID")
		}
	})

	t.Run("Database Error", func(t *testing.T) {
		mockQueries := &mockLoanQueries{
			createErr: errors.New("unique constraint violation"),
		}
		svc := &loanService{queries: mockQueries}
		_, err := svc.CreateLoan(context.Background(), request.CreateLoanRequest{BorrowerId: 123})
		if err == nil {
			t.Error("expected error from database, got nil")
		}
	})
}

func TestInvestLoan_TransactionCalled(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta("SELECT id, borrower_id, principal_amount, rate, roi, agreement_letter_url, state, total_invested, created_at FROM loans WHERE id = ? LIMIT 1")).
		WithArgs("loan-123").
		WillReturnRows(sqlmock.NewRows([]string{
			"id",
			"borrower_id",
			"principal_amount",
			"rate",
			"roi",
			"state",
			"agreement_letter_url",
			"total_invested",
			"created_at",
		}).AddRow("loan-123", 1, 1000.0, 10.0, 12.0, "approved", "http://...", 0.0, time.Now()))

	mock.ExpectQuery(regexp.QuoteMeta("SELECT SUM(amount) AS total_invested FROM investments WHERE loan_id = ?")).
		WithArgs("loan-123").
		WillReturnRows(sqlmock.NewRows([]string{"total_invested"}).AddRow(500.0))

	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO investments (id, loan_id, investor_id, amount) VALUES (?, ?, ?, ?)")).
		WithArgs(sqlmock.AnyArg(), "loan-123", int64(1), 100.0).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectCommit()

	queries := sqlc.New(db)
	svc := &loanService{
		db:        db,
		queries:   queries,
		loanEvent: event.NewLoanEvent(queries),
	}

	_, err = svc.InvestLoan(context.Background(), request.CreateInvestRequest{
		LoanID:     "loan-123",
		InvestorId: 1,
		Amount:     100,
	})

	if err != nil {
		t.Fatalf("InvestLoan returned an error: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
