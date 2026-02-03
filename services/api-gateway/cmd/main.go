package main

import (
	"innoveria-iot/api-gateway/internal/config"
	"innoveria-iot/api-gateway/internal/server"
)

// TODO: Setup logging
func main()  {
	cfg := config.Load()
	server := server.NewServer(cfg)

	server.Run()
}
