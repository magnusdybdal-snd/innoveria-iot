package main

import (
	"innoveria-iot/api-gateway/internal"
	"innoveria-iot/api-gateway/internal/config"
)

func main()  {
	cfg := config.Load()
	server := internal.NewServer(cfg)

	server.Run()
}
