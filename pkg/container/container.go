package container

import (
	"github.com/Pelegrinetti/posweb-user-api/pkg/database"
)

type Container struct {
	Database *database.Database
}

func (c *Container) AddDatabase(db *database.Database) {
	c.Database = db
}

func New() *Container {
	return &Container{}
}
