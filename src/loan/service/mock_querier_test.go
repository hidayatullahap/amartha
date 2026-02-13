package service

import (
	"amartha/generated/sqlc"
	"context"
)

type mockGuardQuerier struct {
	sqlc.Querier
	loan sqlc.Loan
	err  error
}

func (m *mockGuardQuerier) GetLoan(ctx context.Context, id string) (sqlc.Loan, error) {
	return m.loan, m.err
}
