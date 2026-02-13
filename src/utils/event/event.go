package event

import (
	"amartha/generated/sqlc"
	"context"
	"log"
)

type EmailEvent struct {
	LoanID string
}

type LoanEvent struct {
	EmailChan chan EmailEvent
	queries   sqlc.Querier
}

func NewLoanEvent(queries sqlc.Querier) *LoanEvent {
	return &LoanEvent{
		EmailChan: make(chan EmailEvent, 100),
		queries:   queries,
	}
}

func (s *LoanEvent) StartEmailWorker() {
	go func() {
		for event := range s.EmailChan {
			log.Printf("Sending emails to investors for loan: %s", event.LoanID)
			s.processEmails(event.LoanID)
		}
	}()
}

func (s *LoanEvent) processEmails(loanID string) {
	ctx := context.Background()
	emails, err := s.queries.GetInvestorEmailsByLoanID(ctx, loanID)
	if err != nil {
		log.Printf("Error fetching investor emails for loan %s: %v", loanID, err)
		return
	}
	log.Printf("Found %d investors for loan %s. Sending notifications...", len(emails), loanID)

	for _, email := range emails {
		log.Printf("NOTIFICATION: Sending email to %s -> 'Loan %s is now fully funded and invested!'", email, loanID)
	}
}
