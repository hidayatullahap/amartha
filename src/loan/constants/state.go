package constants

import "fmt"

type LoanState int

const (
	StateProposed LoanState = iota
	StateApproved
	StateInvested
	StateDisbursed
)

var stateNames = [...]string{
	StateProposed:  "proposed",
	StateApproved:  "approved",
	StateInvested:  "invested",
	StateDisbursed: "disbursed",
}

var stateMap = map[string]LoanState{
	"proposed":  StateProposed,
	"approved":  StateApproved,
	"invested":  StateInvested,
	"disbursed": StateDisbursed,
}

func (s LoanState) String() string {
	if s < 0 || int(s) >= len(stateNames) {
		return fmt.Sprintf("LoanState(%d)", s)
	}
	return stateNames[s]
}

func (s LoanState) IsValid() bool {
	return s >= StateProposed && s <= StateDisbursed
}

func CanUpdateState(currentState, newState LoanState) bool {
	return newState > currentState && newState.IsValid()
}

func ToLoanState(s string) (LoanState, bool) {
	state, ok := stateMap[s]
	return state, ok
}
