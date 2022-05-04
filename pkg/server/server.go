package server

import (
	"fmt"

	"github.com/Pelegrinetti/posweb-user-api/internal/config"
	"github.com/Pelegrinetti/posweb-user-api/pkg/container"
	"github.com/Pelegrinetti/posweb-user-api/pkg/controllers"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)
type server struct {
	config    config.ServerConfig
	app       *fiber.App
	container *container.Container
}

func (s *server) setupRoutes() {
	s.app.Get("/liveness", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})
	s.app.Get("/readiness", func(c *fiber.Ctx) error {
		ping, pingError := s.container.Database.Ping()

		if pingError != nil {
			logrus.Error("Can't connect with database: ", pingError)

			return c.Status(503).SendString("NOK")
		}

		if ping {
			return c.SendString("OK")
		}

		return c.Status(503).SendString("NOK")
	})

	s.app.Post("/users", controllers.HandleLogin(s.container))
	s.app.Put("/users", controllers.UpdateUser(s.container))
}

func (s *server) Run() {
	s.setupRoutes()

	s.app.Listen(fmt.Sprintf(":%d", s.config.Port))
}

func New(ctn *container.Container) *server {
	app := fiber.New()

	return &server{
		config:    ctn.Config.ServerConfig,
		app:       app,
		container: ctn,
	}
}
