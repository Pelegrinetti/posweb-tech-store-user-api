package database

import (
	"context"
	"time"

	"github.com/Pelegrinetti/posweb-user-api/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Database struct {
	client *mongo.Client
	config config.DatabaseConfig
}

func (db *Database) Ping() (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	err := db.client.Ping(ctx, readpref.Primary())

	if err == nil {
		return true, nil
	}

	return false, err
}

func (db *Database) GetCollection(name string) *mongo.Collection {
	return db.client.Database(db.config.DBName).Collection(name)
}

func New(databaseConfig config.DatabaseConfig) (*Database, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(databaseConfig.Uri))

	db := &Database{
		client: client,
		config: databaseConfig,
	}

	return db, err
}
