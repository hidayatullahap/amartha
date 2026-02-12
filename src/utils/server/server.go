package server

import (
	hello_controller "amartha/src/controllers/hello"
	"fmt"
	"log"
	"os"

	"github.com/google/wire"
	"github.com/labstack/echo/v5"
)

type Server struct {
	echo        *echo.Echo
	controllers *Controllers
}

type Controllers struct {
	HelloController hello_controller.HelloController
}

var RoutesSet = wire.NewSet(wire.Struct(new(Controllers), "*"), NewServer)

func NewServer(
	echo *echo.Echo,
	controllers *Controllers,
) *Server {
	return &Server{
		echo,
		controllers,
	}
}

func (server *Server) Run() {
	server.controllers.HelloController.DecorateRoutes(server.echo)
	log.Fatal(server.echo.Start(fmt.Sprintf(":%s", os.Getenv("PORT"))))
}
