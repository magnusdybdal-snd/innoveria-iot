package main

import (
	"innoveria-iot/api-gateway/internal/config"
	"innoveria-iot/api-gateway/internal/server"
)

func main()  {
	cfg := config.Load()
	server := server.NewServer(cfg)

	server.Run()
}
