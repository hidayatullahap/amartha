package service

import (
	"amartha/src/loan/request"
	"context"
	"errors"
	"testing"
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
