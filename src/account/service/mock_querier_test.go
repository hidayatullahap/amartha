package service

import (
	"amartha/generated/sqlc"
	"context"
)

type mockAccountQuerier struct {
	sqlc.Querier
	user sqlc.GetUserByUsernameRow
	err  error
}

func (m *mockAccountQuerier) GetUserByUsername(ctx context.Context, username string) (sqlc.GetUserByUsernameRow, error) {
	return m.user, m.err
}
