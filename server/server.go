package server

import (
	"github.com/jake-hansen/agora/config"
)

func Init(env string) {
	r := NewRouter(env)
	address := config.GetConfig().GetString("server.address")
	err := r.Run(address)
	if err != nil {
		panic(err)
	}
}