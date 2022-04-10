package main

import (
	"github.com/Pelegrinetti/posweb-user-api/internal/config"
	"github.com/Pelegrinetti/posweb-user-api/pkg/server"
)

func main() {
	cfg := config.New()
	srv := server.New(cfg.ServerConfig)

	srv.Run()
}
