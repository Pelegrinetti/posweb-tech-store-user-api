package models

import (
	"context"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

type Address struct {
	Street       string `bson:"street" json:"street"`
	PostalCode   string `bson:"postal_code" json:"postal_code"`
	Number       int    `bson:"number" json:"number"`
	Neighborhood string `bson:"neighborhood" json:"neighborhood"`
	City         string `bson:"city" json:"city"`
	State        string `bson:"state" json:"state"`
	Country      string `bson:"country" json:"country"`
}

type User struct {
	collection *mongo.Collection
	Id         primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name       string             `bson:"name" json:"name"`
	Email      string             `bson:"email" json:"email"`
	Picture    string             `bson:"picture" json:"picture"`
	Cpf        string             `bson:"cpf" json:"cpf"`
	Cellphone  string             `bson:"cellphone" json:"cellphone"`
	Addresses  []Address          `bson:"addresses" json:"addresses"`
}

func (u *User) Insert() (bool, error) {
	filter := bson.M{"_id": u.Id}
	update := bson.D{{ Key: "$set", Value: u }}
	opts :=  options.Update().SetUpsert(true)

	_, insertError := u.collection.UpdateOne(context.TODO(), filter, update, opts)

	if insertError != nil {
		return false, insertError
	}

	return true, nil
}

func (u *User) FindOne(filter bson.M) (bool, error) {
	result := u.collection.FindOne(context.TODO(), filter)

	decodeError := result.Decode(u)

	if decodeError != nil {
		logrus.Error("Error decoding user on FindOne method: ", decodeError)

		return false, decodeError
	}

	return true, nil
}

func NewUser(collection *mongo.Collection) *User {
	_, indexesError := collection.Indexes().CreateMany(context.Background(), []mongo.IndexModel{
		{
			Keys:    bsonx.Doc{{Key: "cpf", Value: bsonx.Int32(1)}},
			Options: options.Index().SetUnique(true),
		},
		{
			Keys:    bsonx.Doc{{Key: "email", Value: bsonx.Int32(1)}},
			Options: options.Index().SetUnique(true),
		},
	})

	if indexesError != nil {
		logrus.Error("Error creating new indexes: ", indexesError)
	}

	return &User{
		collection: collection,
	}
}
