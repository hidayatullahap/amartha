package service

import (
	"amartha/generated/sqlc"
	"amartha/src/account/response"
	"context"
	"crypto/sha256"
	"encoding/hex"
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
	user, err := s.queries.GetUser(ctx, username)
	if err != nil {
		return nil, err
	}

	h := sha256.New()
	h.Write([]byte(user.ID))
	shaString := hex.EncodeToString(h.Sum(nil))

	return &response.LoginResponse{Token: shaString}, nil
}
