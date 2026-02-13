package service

import (
	"amartha/generated/sqlc"
	"context"
	"database/sql"
	"testing"
)

func TestValidateStatusTransition(t *testing.T) {
	t.Run("Valid Transition: Proposed to Approved", func(t *testing.T) {
		mock := &mockGuardQuerier{
			loan: sqlc.Loan{
				ID:    "loan-1",
				State: sql.NullString{String: "proposed", Valid: true},
			},
		}
		svc := NewGuardLoanService(mock)

		can, loan, err := svc.ValidateStatusTransition(context.Background(), "loan-1", "approved")

		if err != nil || !can {
			t.Errorf("expected transition to be valid, got err: %v, can: %v", err, can)
		}
		if loan.ID != "loan-1" {
			t.Errorf("expected loan-1, got %s", loan.ID)
		}
	})

	t.Run("Invalid Transition: Proposed to Disbursed", func(t *testing.T) {
		mock := &mockGuardQuerier{
			loan: sqlc.Loan{
				ID:    "loan-2",
				State: sql.NullString{String: "proposed", Valid: true},
			},
		}
		svc := NewGuardLoanService(mock)

		can, _, err := svc.ValidateStatusTransition(context.Background(), "loan-2", "disbursed")

		if err == nil {
			t.Error("expected error for invalid state skip, but got nil")
		}
		if can {
			t.Error("expected canUpdate to be false")
		}
	})
}
