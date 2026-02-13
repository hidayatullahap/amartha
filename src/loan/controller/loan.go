package controller

import (
	"amartha/src/loan/constants"
	"amartha/src/loan/request"
	"amartha/src/loan/service"
	"amartha/src/utils/echo/middleware"
	uerror "amartha/src/utils/error"
	"fmt"
	"net/http"

	uecho "amartha/src/utils/echo"

	"github.com/labstack/echo/v5"
)

type LoanController struct {
	svc   service.LoanService
	auth  *middleware.AuthMiddleware
	guard service.GuardLoanService
}

func NewLoanController(
	svc service.LoanService,
	auth *middleware.AuthMiddleware,
	guard service.GuardLoanService,
) LoanController {
	return LoanController{
		svc,
		auth,
		guard,
	}
}

func (r *LoanController) DecorateRoutes(e *echo.Echo) {
	routeGroup := e.Group("/loans", r.auth.Authenticate)
	routeGroup.POST("", func(c *echo.Context) error {
		var body request.CreateLoanRequest

		if err := c.Bind(&body); err != nil {
			code := uerror.GetHttpCodeByError(err)
			return echo.NewHTTPError(code, err.Error())
		}

		user := uecho.GetAuthUser(c)
		body.BorrowerId = user.ID

		data, err := r.svc.CreateLoan(c.Request().Context(), body)
		if err != nil {
			code := uerror.GetHttpCodeByError(err)
			return echo.NewHTTPError(code, err.Error())
		}
		return c.JSON(http.StatusOK, data)
	})

	routeGroup.POST("/:id/approve", func(c *echo.Context) error {
		id := c.Param("id")
		var body request.CreateLoanDetailRequest

		if err := c.Bind(&body); err != nil {
			code := uerror.GetHttpCodeByError(err)
			return echo.NewHTTPError(code, err.Error())
		}
		body.LoanID = id

		_, _, err := r.guard.ValidateStatusTransition(c.Request().Context(), id, constants.StateApproved.String())
		if err != nil {
			code := uerror.GetHttpCodeByError(err)
			return echo.NewHTTPError(code, err.Error())
		}

		data, err := r.svc.ApproveLoan(c.Request().Context(), body)
		if err != nil {
			code := uerror.GetHttpCodeByError(err)
			return echo.NewHTTPError(code, err.Error())
		}
		return c.JSON(http.StatusOK, data)
	}, r.auth.Authorize(constants.AdminRoles))

	routeGroup.POST("/:id/invest", func(c *echo.Context) error {
		id := c.Param("id")
		var body request.CreateInvestRequest

		if err := c.Bind(&body); err != nil {
			code := uerror.GetHttpCodeByError(err)
			return echo.NewHTTPError(code, err.Error())
		}
		body.LoanID = id

		_, _, err := r.guard.ValidateStatusTransition(c.Request().Context(), id, constants.StateInvested.String())
		if err != nil {
			code := uerror.GetHttpCodeByError(err)
			return echo.NewHTTPError(code, err.Error())
		}

		data, err := r.svc.InvestLoan(c.Request().Context(), body)
		if err != nil {
			code := uerror.GetHttpCodeByError(err)
			return echo.NewHTTPError(code, err.Error())
		}
		return c.JSON(http.StatusOK, data)
	})

	routeGroup.POST("/:id/disburse", func(c *echo.Context) error {
		id := c.Param("id")
		_, _, err := r.guard.ValidateStatusTransition(c.Request().Context(), id, constants.StateDisbursed.String())
		if err != nil {
			code := uerror.GetHttpCodeByError(err)
			return echo.NewHTTPError(code, err.Error())
		}

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
