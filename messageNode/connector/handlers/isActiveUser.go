package connectorHandlers

import (
	"SensorManager/messageNode/database"
	"github.com/gofiber/fiber/v2"
)

type IsActiveUserRequest struct {
	User string `json:"user"`
}

func IsActiveUserHandler(c *fiber.Ctx) error {
	p := new(IsActiveUserRequest)
	if err := c.BodyParser(p); err != nil {
		return err
	}

	//receive database instance
	db, ok := c.Locals("db").(*database.Database)
	if !ok {
		return c.Status(fiber.StatusInternalServerError).SendString("Database connection error")
	}
	value := db.Read("active_connection")
	if value == p.User {
		return c.SendStatus(fiber.StatusAccepted)
	} else {
		return c.SendStatus(fiber.StatusForbidden)
	}
}
