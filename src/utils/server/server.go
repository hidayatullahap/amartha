package server

import (
	loanController "amartha/src/loans/controller"
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
	loanController.LoanController
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
	server.controllers.LoanController.DecorateRoutes(server.echo)
	log.Fatal(server.echo.Start(fmt.Sprintf(":%s", os.Getenv("PORT"))))
}
