package echo

import (
	"amartha/src/utils/constants"
	"testing"

	"github.com/labstack/echo/v5"
	"github.com/stretchr/testify/assert"
)

func TestGetAuthUser(t *testing.T) {
	tests := []struct {
		name     string
		setup    func(*echo.Context)
		expected AuthenticatedUser
	}{
		{
			name: "Valid user with ID and role",
			setup: func(c *echo.Context) {
				id := int64(123)
				c.Set(constants.UserId, &id)
				c.Set(constants.UserRole, "admin")
			},
			expected: AuthenticatedUser{
				ID:   123,
				Role: "admin",
			},
		},
		{
			name: "User with role but no ID",
			setup: func(c *echo.Context) {
				c.Set(constants.UserRole, "user")
			},
			expected: AuthenticatedUser{
				ID:   0,
				Role: "user",
			},
		},
		{
			name: "Empty context",
			setup: func(c *echo.Context) {
			},
			expected: AuthenticatedUser{
				ID:   0,
				Role: "",
			},
		},
		{
			name: "Invalid ID type",
			setup: func(c *echo.Context) {
				c.Set(constants.UserId, "not-a-pointer")
				c.Set(constants.UserRole, "admin")
			},
			expected: AuthenticatedUser{
				ID:   0,
				Role: "admin",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &echo.Context{}
			tt.setup(c)
			user := GetAuthUser(c)
			assert.Equal(t, tt.expected, user)
		})
	}
}
