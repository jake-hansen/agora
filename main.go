package main

import (
	"github.com/jake-hansen/agora/server"
)

func main() {
	cleanup := startAPIServer()
	defer cleanup()
}

func startAPIServer() func() {
	apiServer, cleanup, err := server.Build()
	if err != nil {
		panic(err)
	}
	err = apiServer.ListenAndServe()
	if err != nil {
		panic(err)
	}
	return cleanup
}
