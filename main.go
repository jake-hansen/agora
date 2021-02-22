package main

import (
	"github.com/jake-hansen/agora/config"
	"github.com/jake-hansen/agora/database"
	"github.com/jake-hansen/agora/database/loader"
	"github.com/jake-hansen/agora/log"
	"github.com/jake-hansen/agora/server"
	"github.com/spf13/viper"
)

func main() {
	configuration := config.Build()
	logger, logCleanup, err := log.Build()
	if err != nil {
		panic(err)
	}
	db, dbCleanup, err := database.Build(configuration, logger)
	if err != nil {
		panic(err)
	}

	loadData(db, configuration)

	serverCleanup := startAPIServer()

	cleanup := func() {
		serverCleanup()
		dbCleanup()
		logCleanup()
	}

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
