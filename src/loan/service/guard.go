package service

import (
	"amartha/generated/sqlc"
	"amartha/src/loan/constants"
	uerror "amartha/src/utils/error"
	"context"
)

type GuardLoanService interface {
	ValidateStatusTransition(ctx context.Context, loanId string, targetStatus string) (bool, *sqlc.Loan, error)
}

type guardLoanService struct {
	queries sqlc.Querier
}

func NewGuardLoanService(queries sqlc.Querier) GuardLoanService {
	return guardLoanService{queries}
}

func (s guardLoanService) ValidateStatusTransition(ctx context.Context, loanId string, targetStatus string) (bool, *sqlc.Loan, error) {
	var result bool
	loan, err := s.queries.GetLoan(ctx, loanId)
	if err != nil {
		return result, nil, err
	}

	prev, _ := constants.ToLoanState(loan.State.String)
	next, _ := constants.ToLoanState(targetStatus)
	canUpdate := constants.CanUpdateState(prev, next)
	if !canUpdate {
		return result, &loan, uerror.ErrInvalidLoanStateUpdate
	}
	return canUpdate, &loan, nil
}
