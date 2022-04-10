package server

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type server struct {
	port int
	app  *fiber.App
}

func (s *server) setupRoutes() {
	s.app.Get("/*", func(c *fiber.Ctx) error {
		return c.SendString("Olá mundo")
	})
}

func (s *server) Run() {
	s.setupRoutes()

	s.app.Listen(fmt.Sprintf(":%d", s.port))
}

func New(port int) *server {
	app := fiber.New()

	return &server{
		port,
		app,
	}
}
