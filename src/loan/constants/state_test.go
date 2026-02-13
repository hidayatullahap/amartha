package constants

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToLoanState(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected LoanState
		ok       bool
	}{
		{
			name:     "Valid proposed state",
			input:    "proposed",
			expected: StateProposed,
			ok:       true,
		},
		{
			name:     "Valid approved state",
			input:    "approved",
			expected: StateApproved,
			ok:       true,
		},
		{
			name:     "Valid invested state",
			input:    "invested",
			expected: StateInvested,
			ok:       true,
		},
		{
			name:     "Valid disbursed state",
			input:    "disbursed",
			expected: StateDisbursed,
			ok:       true,
		},
		{
			name:     "Invalid state",
			input:    "unknown",
			expected: 0,
			ok:       false,
		},
		{
			name:     "Empty string",
			input:    "",
			expected: 0,
			ok:       false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			state, ok := ToLoanState(tt.input)
			assert.Equal(t, tt.expected, state)
			assert.Equal(t, tt.ok, ok)
		})
	}
}
func TestCanUpdateState(t *testing.T) {
	tests := []struct {
		name         string
		currentState LoanState
		newState     LoanState
		expected     bool
	}{
		{
			name:         "Valid transition from proposed to approved",
			currentState: StateProposed,
			newState:     StateApproved,
			expected:     true,
		},
		{
			name:         "Valid transition from approved to invested",
			currentState: StateApproved,
			newState:     StateInvested,
			expected:     true,
		},
		{
			name:         "Valid transition from invested to disbursed",
			currentState: StateInvested,
			newState:     StateDisbursed,
			expected:     true,
		},
		{
			name:         "Valid transition from proposed to disbursed",
			currentState: StateProposed,
			newState:     StateDisbursed,
			expected:     false,
		},
		{
			name:         "Invalid transition backwards",
			currentState: StateApproved,
			newState:     StateProposed,
			expected:     false,
		},
		{
			name:         "Valid transition same state",
			currentState: StateApproved,
			newState:     StateApproved,
			expected:     true,
		},
		{
			name:         "Invalid transition with invalid new state",
			currentState: StateProposed,
			newState:     LoanState(999),
			expected:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CanUpdateState(tt.currentState, tt.newState)
			assert.Equal(t, tt.expected, result)
		})
	}
}
func TestString(t *testing.T) {
	tests := []struct {
		name     string
		state    LoanState
		expected string
	}{
		{
			name:     "StateProposed string representation",
			state:    StateProposed,
			expected: "proposed",
		},
		{
			name:     "StateApproved string representation",
			state:    StateApproved,
			expected: "approved",
		},
		{
			name:     "StateInvested string representation",
			state:    StateInvested,
			expected: "invested",
		},
		{
			name:     "StateDisbursed string representation",
			state:    StateDisbursed,
			expected: "disbursed",
		},
		{
			name:     "Invalid negative state",
			state:    LoanState(-1),
			expected: "LoanState(-1)",
		},
		{
			name:     "Invalid state out of range",
			state:    LoanState(999),
			expected: "LoanState(999)",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.state.String()
			assert.Equal(t, tt.expected, result)
		})
	}
}
