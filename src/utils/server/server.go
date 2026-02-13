package server

import (
	accountController "amartha/src/account/controller"
	loanController "amartha/src/loan/controller"
	"amartha/src/utils/config"
	"fmt"
	"log"

	"github.com/google/wire"
	"github.com/labstack/echo/v5"
)

type Server struct {
	echo        *echo.Echo
	controllers *Controllers
	config      *config.Config
}

type Controllers struct {
	loanController.LoanController
	accountController.AccountController
}

var RoutesSet = wire.NewSet(wire.Struct(new(Controllers), "*"), NewServer)

func NewServer(
	echo *echo.Echo,
	controllers *Controllers,
	config *config.Config,
) *Server {
	return &Server{
		echo,
		controllers,
		config,
	}
}

func (server *Server) Run() {
	server.controllers.LoanController.DecorateRoutes(server.echo)
	log.Fatal(server.echo.Start(fmt.Sprintf(":%d", server.config.Server.Port)))
}
