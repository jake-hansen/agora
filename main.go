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
	logger, logCleanup, err := log.Build(configuration)
	if err != nil {
		panic(err)
	}
	db, dbCleanup, err := database.Build(configuration, logger)
	if err != nil {
		panic(err)
	}

	loadData(db, configuration)

	startAPIServer(db, configuration, logger)

	cleanup := func() {
		dbCleanup()
		logCleanup()
	}

	defer cleanup()
}

func startAPIServer(db *database.Manager, v *viper.Viper, log *log.Log) {
	apiServer, err := server.Build(db, v, log)
	if err != nil {
		panic(err)
	}
	err = apiServer.ListenAndServe()
	if err != nil {
		panic(err)
	}
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
