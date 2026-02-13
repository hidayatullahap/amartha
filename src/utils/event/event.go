package event

import "log"

type EmailEvent struct {
	LoanID string
}

type LoanEvent struct {
	EmailChan chan EmailEvent
}

func NewLoanEvent() *LoanEvent {
	return &LoanEvent{
		EmailChan: make(chan EmailEvent, 100),
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
	log.Println("Process Email", loanID)
}
