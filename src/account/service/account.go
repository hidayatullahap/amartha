package service

import (
	"amartha/generated/sqlc"
	"amartha/src/account/response"
	"amartha/src/utils/crypto"
	"context"
)

type AccountService interface {
	Login(ctx context.Context, username string) (*response.LoginResponse, error)
}

type accountService struct {
	queries *sqlc.Queries
}

func NewAccountService(queries *sqlc.Queries) AccountService {
	return accountService{queries}
}

func (s accountService) Login(ctx context.Context, username string) (*response.LoginResponse, error) {
	user, err := s.queries.GetUserByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	jwt, err := crypto.CreateToken(user.ID)
	if err != nil {
		return nil, err
	}

	return &response.LoginResponse{Token: jwt}, nil
}
