package controller

import (
	"amartha/src/account/request"
	"amartha/src/account/service"
	"net/http"

	uerror "amartha/src/utils/error"

	"github.com/labstack/echo/v5"
)

type AccountController struct {
	svc service.AccountService
}

func NewAccountController(svc service.AccountService) AccountController {
	return AccountController{
		svc,
	}
}

func (r *AccountController) DecorateRoutes(e *echo.Echo) {
	routeGroup := e.Group("/account")
	routeGroup.POST("/login", func(c *echo.Context) error {
		body := new(request.LogiRequest)

		if err := c.Bind(body); err != nil {
			code := uerror.GetHttpCodeByError(err)
			return echo.NewHTTPError(code, err.Error())
		}

		data, err := r.svc.Login(c.Request().Context(), body.Username)
		if err != nil {
			code := uerror.GetHttpCodeByError(err)
			return echo.NewHTTPError(code, err.Error())
		}
		return c.JSON(http.StatusOK, data)
	})
}
