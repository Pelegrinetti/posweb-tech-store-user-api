package main

import (
	"github.com/Pelegrinetti/posweb-user-api/pkg/server"
)

func main() {
	s := server.New(3001)

	s.Run()
}
