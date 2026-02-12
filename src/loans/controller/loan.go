package controller

import (
	"net/http"

	"github.com/labstack/echo/v5"
)

type LoanController struct {
}

func NewLoanController() LoanController {
	return LoanController{}
}

func (r *LoanController) DecorateRoutes(e *echo.Echo) {
	routeGroup := e.Group("/loans")
	routeGroup.GET("", func(c *echo.Context) error {
		return c.JSON(http.StatusOK, "Loan endpoint")
	})
}
