package main

import (
	"github.com/jake-hansen/agora/config"
	"github.com/jake-hansen/agora/database"
	"github.com/jake-hansen/agora/database/loader"
	"github.com/jake-hansen/agora/server"
	"github.com/spf13/viper"
)

func main() {
	configuration := config.Build()
	db, dbCleanup, err := database.Build()
	if err != nil {
		panic(err)
	}

	loadData(db, configuration)

	cleanup := startAPIServer()
	defer cleanup()
	defer dbCleanup()
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

func loadData(db *database.Manager, v *viper.Viper) {
	loader, err := loader.Build(db, v)
	if err != nil {
		panic(err)
	}
	err = loader.Load()
	if err != nil {
		panic(err)
	}
}
