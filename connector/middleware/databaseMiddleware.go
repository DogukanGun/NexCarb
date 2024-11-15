package connectormiddleware

import (
	"SensorManager/database"
	"github.com/gofiber/fiber/v2"
)

func DatabaseMiddleware(db *database.Database) fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Locals("db", db)
		return c.Next()
	}
}
