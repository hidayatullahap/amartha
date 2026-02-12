package hello_controller

import (
	"net/http"

	"github.com/labstack/echo/v5"
)

type HelloController struct {
}

func NewHelloController() HelloController {
	return HelloController{}
}

func (r *HelloController) DecorateRoutes(e *echo.Echo) {
	routeGroup := e.Group("/hello")
	routeGroup.GET("", func(c *echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
}
