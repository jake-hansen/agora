package server

import (
	"github.com/jake-hansen/agora/router"
)

// Config contains the parameters for configuring a Server.
type Config struct {
	Address string
}

// Server contains the Config and Router for serving requests.
type Server struct {
	config *Config
	router *router.Router
}

// ListenAndServe begins serving requests.
func (s *Server) ListenAndServe() error {
	err := s.router.Run(s.config.Address)
	return err
}

// New returns a new Server configured with the given Config and Router.
func New(cfg Config, router *router.Router) *Server {
	s := &Server{
		config: &cfg,
		router: router,
	}
	return s
}
