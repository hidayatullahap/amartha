package middleware

import (
	"amartha/generated/sqlc"
	"log"

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
		log.Println(authHeader)

		return next(c)
	}
}
