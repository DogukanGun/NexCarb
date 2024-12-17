package connectorHandlers

import (
	"SensorManager/messageNode/database"
	"github.com/gofiber/fiber/v2"
)

type ReceivedMessage struct {
	Sender  string `json:"sender"`
	Message string `json:"message"`
}

func ReceiveMessageHandler(c *fiber.Ctx) error {
	p := new(SendMessageRequest)
	if err := c.BodyParser(p); err != nil {
		return err
	}
	db, ok := c.Locals("db").(*database.Database)
	if !ok {
		return c.Status(fiber.StatusInternalServerError).SendString("Database connection error")
	}
	//TODO encrypt the message and save database
	err := db.Write(database.MESSAGE, "")
	if err != nil {

	}
	return nil
}
