package main

import (
	"amartha/src/wire"
)

func main() {
	server, cleanup, err := wire.Initialize()
	if err != nil {
		panic(err)
	}

	defer func() {
		cleanup()
	}()

	server.Run()
}
