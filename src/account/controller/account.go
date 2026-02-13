package controller

import (
	"amartha/src/account/service"
	"net/http"

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
		r.svc.Login()
		return c.JSON(http.StatusOK, "[TODO] Login")
	})
}
