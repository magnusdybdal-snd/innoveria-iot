package server

import (
	"innoveria-iot/api-gateway/internal/config"
	"innoveria-iot/api-gateway/internal/routing"
	"net/http"
)

type Server struct {
	cfg *config.Config
}

func NewServer(cfg *config.Config) *Server {
	return &Server{
		cfg: cfg,
		// logger
	}
}

func (s *Server) Run() {
	mux := routing.NewRouter()

	server :=  &http.Server{
		Addr: s.cfg.Addr,
		Handler: mux,
	}

	// TODO: graceful shutdown?
	server.ListenAndServe()
}

