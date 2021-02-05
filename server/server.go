package server

import (
	"github.com/jake-hansen/agora/router"
)

type Config struct {
	Address		string
}

type Server struct {
	config		*Config
	router		*router.Router
}

func (s *Server) ListenAndServe() error {
	err := s.router.Run(s.config.Address)
	return err
}

func New(cfg Config, router *router.Router) *Server  {
	s := &Server{
		config: &cfg,
		router: router,
	}
	return s
}