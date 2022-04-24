package controllers

import (
	"encoding/json"

	"github.com/Pelegrinetti/posweb-user-api/pkg/database"
	"github.com/Pelegrinetti/posweb-user-api/pkg/models"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateUser(db *database.Database) fiber.Handler {
	return func(c *fiber.Ctx) error {
		usersCollection := db.GetCollection("users")

		user := models.NewUser(usersCollection)

		unmarshalError := json.Unmarshal(c.Body(), user)

		if unmarshalError != nil {
			logrus.Error(unmarshalError)

			return c.SendStatus(500)
		}

		user.Id = primitive.NewObjectID()

		_, insertError := user.Insert()

		if insertError != nil {
			logrus.Error("Insert user error: ", insertError)

			return c.SendStatus(400)
		}

		parsedUser, parsingError := json.Marshal(user)

		if parsingError != nil {
			logrus.Error("Parsing user to JSON: ", parsingError)

			return c.SendStatus(500)
		}

		return c.Status(201).Send(parsedUser)
	}
}

func UpdateUser(db *database.Database) fiber.Handler {
	return func(c *fiber.Ctx) error {
		usersCollection := db.GetCollection("users")

		user := models.NewUser(usersCollection)

		_, findError := user.FindOne(c.Params("id"))

		if findError != nil {
			logrus.Error("Can't find user: ", findError)
		}

		parsedUser, parsingError := json.Marshal(user)

		if parsingError != nil {
			logrus.Error("Parsing user to JSON: ", parsingError)

			return c.SendStatus(500)
		}

		return c.Status(201).Send(parsedUser)
	}
}
