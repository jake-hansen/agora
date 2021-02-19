package main

import (
	"github.com/jake-hansen/agora/database/loader"
	"github.com/jake-hansen/agora/server"
)

func main() {
	loadData()

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

func loadData() {
	loader, cleanup, err := loader.Build()
	if err != nil {
		panic(err)
	}
	err = loader.Load()
	if err != nil {
		panic(err)
	}
	cleanup()
}
