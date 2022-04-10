package container

import (
	"github.com/Pelegrinetti/posweb-user-api/pkg/database"
)

type container struct {
	database *database.Database
}

func (c *container) AddDatabase(db *database.Database) {
	c.database = db
}

func New() *container {
	return &container{}
}
