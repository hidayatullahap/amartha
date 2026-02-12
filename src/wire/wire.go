//go:build wireinject
// +build wireinject

package wire

import (
	hello_controller "amartha/src/controllers/hello"

	"amartha/src/utils/echo"
	"amartha/src/utils/server"

	"github.com/google/wire"
	_ "github.com/joho/godotenv/autoload"
)

func Initialize() (*server.Server, func(), error) {
	panic(
		wire.Build(
			hello_controller.NewHelloController,
			echo.NewEcho,
			server.RoutesSet,
		),
	)
}
