package container

import (
	"github.com/Pelegrinetti/posweb-user-api/internal/config"
	"github.com/Pelegrinetti/posweb-user-api/pkg/database"
)

type Container struct {
	Database *database.Database
	Config   *config.Config
}

func (c *Container) AddDatabase(db *database.Database) {
	c.Database = db
}

func (c *Container) AddConfig(config *config.Config) {
	c.Config = config
}

func New() *Container {
	return &Container{}
}
