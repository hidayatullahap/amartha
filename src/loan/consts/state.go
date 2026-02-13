package loans

type LoanState int

const (
	StateProposed LoanState = iota
	StateApproved
	StateInvested
	StateDisbursed
)

func (s LoanState) String() string {
	return [...]string{"proposed", "approved", "invested", "disbursed"}[s]
}

func CanUpdateState(currentState LoanState, newState LoanState) bool {
	return newState > currentState
}
