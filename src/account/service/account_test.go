package service

import (
	"amartha/generated/sqlc"
	"context"
	"errors"
	"testing"
)

func TestLogin_Success(t *testing.T) {
	mock := &mockAccountQuerier{
		user: sqlc.GetUserByUsernameRow{
			ID:       1,
			Username: "testuser",
		},
	}

	svc := NewAccountService(mock)
	resp, err := svc.Login(context.Background(), "testuser")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if resp.Token == "" {
		t.Error("expected a token, but got an empty string")
	}
}

func TestLogin_UserNotFound(t *testing.T) {
	mock := &mockAccountQuerier{
		err: context.DeadlineExceeded,
	}

	svc := NewAccountService(mock)
	resp, err := svc.Login(context.Background(), "unknown")
	if err == nil {
		t.Error("expected error for unknown user, but got nil")
	}
	if resp != nil {
		t.Error("expected response to be nil on error")
	}
}

func TestLogin_CreateTokenError(t *testing.T) {
	old := createToken
	defer func() { createToken = old }()

	createToken = func(id string) (string, error) {
		return "", errors.New("mock token error")
	}

	mockDB := &mockAccountQuerier{
		user: sqlc.GetUserByUsernameRow{ID: 1, Username: "test"},
	}
	svc := NewAccountService(mockDB)
	_, err := svc.Login(context.Background(), "test")

	if err == nil || err.Error() != "mock token error" {
		t.Errorf("expected mock token error, got %v", err)
	}
}
