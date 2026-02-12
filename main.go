package main

import (
	"amartha/src/wire"
	"net/http"
)

type roundTripper struct {
	rt http.RoundTripper
}

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
