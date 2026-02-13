package service

import (
	"amartha/generated/sqlc"
	"context"
)

type GuardLoanService interface {
	ValidateStatusTransition(ctx context.Context, loanId string, targetStatus string) (bool, error)
}

type guardLoanService struct {
	queries *sqlc.Queries
}

func NewGuardLoanService(queries *sqlc.Queries) GuardLoanService {
	return guardLoanService{queries}
}

func (g guardLoanService) ValidateStatusTransition(ctx context.Context, loanId string, targetStatus string) (bool, error) {
	panic("unimplemented")
}
