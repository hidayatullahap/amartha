//go:build wireinject
// +build wireinject

package wire

import (
	accountController "amartha/src/account/controller"
	accountService "amartha/src/account/service"
	loanController "amartha/src/loan/controller"
	loanService "amartha/src/loan/service"

	"amartha/src/utils/config"
	"amartha/src/utils/database"
	"amartha/src/utils/echo"
	"amartha/src/utils/echo/middleware"
	"amartha/src/utils/server"

	"github.com/google/wire"
	_ "github.com/joho/godotenv/autoload"
)

func Initialize() (*server.Server, func(), error) {
	panic(
		wire.Build(
			middleware.NewAuthMiddleware,
			database.NewQueries,
			database.NewDBConnection,
			config.LoadConfig,
			loanService.NewLoanService,
			accountService.NewAccountService,
			loanController.NewLoanController,
			accountController.NewAccountController,
			echo.NewEcho,
			server.RoutesSet,
		),
	)
}
