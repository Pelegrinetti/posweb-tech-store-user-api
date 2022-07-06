package controllers

import (
	"encoding/json"

	"github.com/Pelegrinetti/posweb-user-api/pkg/container"
	"github.com/Pelegrinetti/posweb-user-api/pkg/models"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func HandleLogin(ctn *container.Container) fiber.Handler {
	db := ctn.Database

	return func(c *fiber.Ctx) error {
		usersCollection := db.GetCollection("users")
		user := models.NewUser(usersCollection)

		unmarshalError := json.Unmarshal(c.Body(), &user)

		if unmarshalError != nil {
			logrus.Error(unmarshalError)

			return c.SendStatus(500)
		}

		if found, _ := user.FindOne(bson.M{
			"email": user.Email,
		}); found {
			jsonOutput, jsonOutputError := json.Marshal(user)

			if jsonOutputError != nil {
				logrus.Error("Can't parse json: ", jsonOutputError)

				return c.SendStatus(500)
			}

			return c.Status(200).Send(jsonOutput)
		}

		user.Id = primitive.NewObjectID()

		_, insertError := user.Insert()

		if insertError != nil {
			logrus.Error("Insert user error: ", insertError)

			return c.SendStatus(400)
		}

		jsonOutput, jsonOutputError := json.Marshal(user)

		if jsonOutputError != nil {
			logrus.Error("Can't parse json: ", jsonOutputError)

			return c.SendStatus(500)
		}

		return c.Status(201).Send(jsonOutput)
	}
}

func UpdateUser(ctn *container.Container) fiber.Handler {
	db := ctn.Database

	return func(c *fiber.Ctx) error {
		usersCollection := db.GetCollection("users")

		user := models.NewUser(usersCollection)

		objectId, err := primitive.ObjectIDFromHex(c.Params("id"))
		if err != nil {
			logrus.Error("Invalid ID: ", err)

			return c.SendStatus(400)
		}

		_, findError := user.FindOne(bson.M{
			"_id": objectId,
		})

		if findError != nil {
			logrus.Error("Can't find user: ", findError)
		}

		unmarshalError := json.Unmarshal(c.Body(), &user)

		if unmarshalError != nil {
			logrus.Error(unmarshalError)

			return c.SendStatus(500)
		}

		_, insertError := user.Insert()

		if insertError != nil {
			logrus.Error("Insert user error: ", insertError)

			return c.SendStatus(500)
		}

		parsedUser, parsingError := json.Marshal(user)

		if parsingError != nil {
			logrus.Error("Parsing user to JSON: ", parsingError)

			return c.SendStatus(500)
		}

		return c.Status(200).Send(parsedUser)
	}
}
