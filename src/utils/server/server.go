package server

import (
	accountController "amartha/src/account/controller"
	loanController "amartha/src/loan/controller"
	"amartha/src/utils/config"
	"amartha/src/utils/event"
	"fmt"
	"log"

	"github.com/google/wire"
	"github.com/labstack/echo/v5"
)

type Server struct {
	echo        *echo.Echo
	controllers *Controllers
	config      *config.Config
	loanEvent   *event.LoanEvent
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
	loanEvent *event.LoanEvent,
) *Server {
	return &Server{
		echo,
		controllers,
		config,
		loanEvent,
	}
}

func (server *Server) Run() {
	server.loanEvent.StartEmailWorker()
	server.controllers.AccountController.DecorateRoutes(server.echo)
	server.controllers.LoanController.DecorateRoutes(server.echo)
	log.Fatal(server.echo.Start(fmt.Sprintf(":%d", server.config.Server.Port)))
}
