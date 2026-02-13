package controller

import (
	"amartha/src/loan/constants"
	"amartha/src/loan/service"
	"amartha/src/utils/echo/middleware"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v5"
)

type LoanController struct {
	svc  service.LoanService
	auth *middleware.AuthMiddleware
}

func NewLoanController(svc service.LoanService, auth *middleware.AuthMiddleware) LoanController {
	return LoanController{
		svc,
		auth,
	}
}

func (r *LoanController) DecorateRoutes(e *echo.Echo) {
	routeGroup := e.Group("/loans", r.auth.Authenticate)
	routeGroup.POST("", func(c *echo.Context) error {
		r.svc.CreateLoan()
		return c.JSON(http.StatusOK, "[TODO] Create a new loan state: proposed")
	})
	routeGroup.POST("/:id/approve", func(c *echo.Context) error {
		id := c.Param("id")
		r.svc.ApproveLoan()
		return c.JSON(http.StatusOK, fmt.Sprintf("[TODO] Approve loan id: %s", id))
	}, r.auth.Authorize(constants.AdminRoles))
	routeGroup.POST("/:id/invest", func(c *echo.Context) error {
		id := c.Param("id")
		r.svc.InvestLoan()
		return c.JSON(http.StatusOK, fmt.Sprintf("[TODO] Invest loan id: %s", id))
	})
	routeGroup.POST("/:id/disburse", func(c *echo.Context) error {
		id := c.Param("id")
		r.svc.DisburseLoan()
		return c.JSON(http.StatusOK, fmt.Sprintf("[TODO] Disburse loan id: %s", id))
	})
	routeGroup.GET("/:id", func(c *echo.Context) error {
		id := c.Param("id")
		r.svc.GetLoan()
		return c.JSON(http.StatusOK, fmt.Sprintf("[TODO] Loan details for id: %s", id))
	})
	routeGroup.GET("", func(c *echo.Context) error {
		r.svc.GetLoans()
		return c.JSON(http.StatusOK, "[TODO] Loan list")
	})
}
