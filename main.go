package main

import (
	"github.com/jake-hansen/agora/server"
)

func main() {
	startAPIServer()
}

func startAPIServer() {
	apiServer, err := server.Build()
	if err != nil {
		panic(err)
	}
	err = apiServer.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
