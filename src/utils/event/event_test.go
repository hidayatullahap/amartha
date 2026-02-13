package event

import (
	"amartha/generated/sqlc"
	"context"
	"testing"
	"time"
)

type mockQuerier struct {
	sqlc.Querier
	calledWith string
}

func (m *mockQuerier) GetInvestorEmailsByLoanID(ctx context.Context, loanID string) ([]string, error) {
	m.calledWith = loanID
	return []string{"investor@example.com"}, nil
}

func TestEmailWorker(t *testing.T) {
	mock := &mockQuerier{}
	s := &LoanEvent{
		EmailChan: make(chan EmailEvent, 1),
		queries:   mock,
	}

	s.StartEmailWorker()

	s.EmailChan <- EmailEvent{LoanID: "LOAN-100"}
	time.Sleep(50 * time.Millisecond)

	if mock.calledWith != "LOAN-100" {
		t.Errorf("Expected mock to be called with LOAN-100, got %s", mock.calledWith)
	}
}
