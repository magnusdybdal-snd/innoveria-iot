package internal

import (
	"innoveria-iot/api-gateway/internal/config"
	"net/http"
)

type Server struct {
	cfg *config.Config
}

func NewServer(cfg *config.Config) *Server {
	return &Server{
		cfg: cfg,
	}
}

func (s *Server) Run() {
	mux := NewRouter()
	server :=  &http.Server{
		Addr: s.cfg.Addr,
		Handler: mux,
	}

	server.ListenAndServe()
}

