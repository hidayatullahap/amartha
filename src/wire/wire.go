//go:build wireinject
// +build wireinject

package wire

import (
	loanController "amartha/src/loan/controller"
	loanService "amartha/src/loan/service"

	"amartha/src/utils/config"
	"amartha/src/utils/echo"
	"amartha/src/utils/server"

	"github.com/google/wire"
	_ "github.com/joho/godotenv/autoload"
)

func Initialize() (*server.Server, func(), error) {
	panic(
		wire.Build(
			config.LoadConfig,
			loanService.NewLoanService,
			loanController.NewLoanController,
			echo.NewEcho,
			server.RoutesSet,
		),
	)
}
