package controller

import (
	"fmt"
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
	routeGroup.POST("", func(c *echo.Context) error {
		return c.JSON(http.StatusOK, "[TODO] Create a new loan state: proposed")
	})
	routeGroup.POST("/:id/approve", func(c *echo.Context) error {
		id := c.Param("id")
		return c.JSON(http.StatusOK, fmt.Sprintf("[TODO] Approve loan id: %s", id))
	})
	routeGroup.POST("/:id/invest", func(c *echo.Context) error {
		id := c.Param("id")
		return c.JSON(http.StatusOK, fmt.Sprintf("[TODO] Invest loan id: %s", id))
	})
	routeGroup.POST("/:id/disburse", func(c *echo.Context) error {
		id := c.Param("id")
		return c.JSON(http.StatusOK, fmt.Sprintf("[TODO] Disburse loan id: %s", id))
	})
	routeGroup.GET("/:id", func(c *echo.Context) error {
		id := c.Param("id")
		return c.JSON(http.StatusOK, fmt.Sprintf("[TODO] Loan details for id: %s", id))
	})
}
