package middleware

import (
	"amartha/generated/sqlc"
	"amartha/src/utils/constants"
	"amartha/src/utils/crypto"

	"github.com/labstack/echo/v5"
)

type AuthMiddleware struct {
	queries *sqlc.Queries
}

func NewAuthMiddleware(queries *sqlc.Queries) *AuthMiddleware {
	return &AuthMiddleware{queries}
}

func (m *AuthMiddleware) Authenticate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c *echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return echo.ErrUnauthorized
		}

		userId, err := crypto.VerifyToken(authHeader)
		if err != nil {
			return echo.ErrUnauthorized
		}

		user, err := m.queries.GetUserById(c.Request().Context(), userId)
		if err != nil {
			return echo.ErrUnauthorized
		}

		c.Set(constants.UserId, userId)
		c.Set(constants.UserRole, user.Role)

		return next(c)
	}
}
