//go:build wireinject
// +build wireinject

package wire

import (
	loanController "amartha/src/loans/controller"

	"amartha/src/utils/echo"
	"amartha/src/utils/server"

	"github.com/google/wire"
	_ "github.com/joho/godotenv/autoload"
)

func Initialize() (*server.Server, func(), error) {
	panic(
		wire.Build(
			loanController.NewLoanController,
			echo.NewEcho,
			server.RoutesSet,
		),
	)
}
