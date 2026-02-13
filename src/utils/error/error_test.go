package error_test

import (
	uerror "amartha/src/utils/error"
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetHttpCodeByError(t *testing.T) {
	tests := []struct {
		name     string
		err      error
		expected int
	}{
		{
			name:     "General error",
			err:      uerror.ErrGeneral,
			expected: http.StatusInternalServerError,
		},
		{
			name:     "Data not found error",
			err:      uerror.ErrDataNotFound,
			expected: http.StatusNotFound,
		},
		{
			name:     "Unauthorized error",
			err:      uerror.ErrUnauthorized,
			expected: http.StatusUnauthorized,
		},
		{
			name:     "Undefined error",
			err:      errors.New("unknown error"),
			expected: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			code := uerror.GetHttpCodeByError(tt.err)
			assert.Equal(t, tt.expected, code)
		})
	}
}
