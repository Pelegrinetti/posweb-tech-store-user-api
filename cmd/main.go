package main

import (
	"github.com/Pelegrinetti/posweb-user-api/internal/config"
	"github.com/Pelegrinetti/posweb-user-api/pkg/container"
	"github.com/Pelegrinetti/posweb-user-api/pkg/database"
	"github.com/Pelegrinetti/posweb-user-api/pkg/server"
	"github.com/sirupsen/logrus"
)

func checkError(err error) {
	if err != nil {
		logrus.Fatal(err)
	}
}

func main() {
	cfg := config.New()

	db, dbError := database.New(cfg.DatabaseConfig)

	ctn := container.New()

	ctn.AddDatabase(db)

	srv := server.New(ctn, cfg.ServerConfig)

	checkError(dbError)

	srv.Run()
}
