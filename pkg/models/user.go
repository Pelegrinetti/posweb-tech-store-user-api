package models

import (
	"context"
	"errors"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/paemuri/brdoc"
	"github.com/sirupsen/logrus"
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

func validateCpf(value interface{}) error {
	cpf := value.(string)

	if !brdoc.IsCPF(cpf) {
		return errors.New("must be a valid cpf")
	}

	return nil
}

func (u *User) Validate() error {
	return validation.ValidateStruct(u,
		validation.Field(&u.Id, validation.Required),
		validation.Field(&u.Name, validation.Required, validation.Length(3, 80)),
		validation.Field(&u.Email, validation.Required, is.Email),
		validation.Field(&u.Cellphone, validation.NilOrNotEmpty),
		validation.Field(&u.Picture, validation.NilOrNotEmpty),
		validation.Field(&u.Cpf, validation.NilOrNotEmpty, validation.By(validateCpf)),
		validation.Field(&u.Addresses, validation.NilOrNotEmpty),
	)
}

func (u *User) Insert() (bool, error) {
	validationError := u.Validate()

	if validationError != nil {
		return false, validationError
	}

	_, insertError := u.collection.InsertOne(context.TODO(), u)

	if insertError != nil {
		return false, insertError
	}

	return true, nil
}

func New(collection *mongo.Collection) *User {
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
