package database

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type DatabaseConfig struct {
	Uri string `mapstructure:"MONGODB_URI"`
}

type Database struct {
	client *mongo.Client
	config DatabaseConfig
}

func (db *Database) Ping() (bool, error) {
	err := db.client.Ping(context.TODO(), readpref.Primary())

	if err == nil {
		return true, nil
	}

	return false, err
}

func New(databaseConfig DatabaseConfig) (*Database, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(databaseConfig.Uri))

	db := &Database{
		client: client,
		config: databaseConfig,
	}

	return db, err
}
