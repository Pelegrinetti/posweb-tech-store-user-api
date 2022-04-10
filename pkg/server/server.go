package server

import (
	"fmt"

	"github.com/Pelegrinetti/posweb-user-api/pkg/container"
	"github.com/gofiber/fiber/v2"
)

type ServerConfig struct {
	Port int `mapstructure:"PORT"`
}
type server struct {
	config    ServerConfig
	app       *fiber.App
	container *container.Container
}

func (s *server) setupRoutes() {
	s.app.Get("/liveness", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})
	s.app.Get("/readiness", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})
}

func (s *server) Run() {
	s.setupRoutes()

	s.app.Listen(fmt.Sprintf(":%d", s.config.Port))
}

func New(ctn *container.Container, serverConfig ServerConfig) *server {
	app := fiber.New()

	return &server{
		config:    serverConfig,
		app:       app,
		container: ctn,
	}
}
