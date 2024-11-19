package connectorMiddleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
)

func DatabaseMiddleware(db *pgx.Conn) fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Locals("db", db)
		return c.Next()
	}
}
