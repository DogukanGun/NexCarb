package connectorHandlers

import (
	"SensorManager/database"
	"github.com/gofiber/fiber/v2"
)

type ConnectionRequest struct {
	Message string `json:"name"`
}

func ConnectRequestHandler(c *fiber.Ctx) error {
	p := new(ConnectionRequest)
	if err := c.BodyParser(p); err != nil {
		return err
	}

	//receive database instance
	db, ok := c.Locals("db").(*database.Database)
	if !ok {
		return c.Status(fiber.StatusInternalServerError).SendString("Database connection error")
	}
	value := db.Read("active_connection")
	if value == "" {
		return c.Status(fiber.StatusForbidden).SendString("Active connection is not possible")
	} else {
		return c.Status(fiber.StatusAccepted).SendString("Please pay the fee first")
	}
}
