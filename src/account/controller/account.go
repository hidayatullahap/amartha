package controller

import (
	"net/http"

	"github.com/labstack/echo/v5"
)

type AccountController struct {
}

func NewAccountController() AccountController {
	return AccountController{}
}

func (r *AccountController) DecorateRoutes(e *echo.Echo) {
	routeGroup := e.Group("/account")
	routeGroup.POST("/login", func(c *echo.Context) error {
		return c.JSON(http.StatusOK, "[TODO] Login")
	})
}
