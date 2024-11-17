package connectorHandlers

import (
	"SensorManager/messageNode/database"
	"SensorManager/utils"
)

type ConnectionRequest struct {
	SenderWallet string `json:"sender_wallet"`
	Message      string `json:"name"`
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
		utils.RunWithHandlingError(db.Write("user_"+p.SenderWallet+"_status", "payment_pending"))
		return c.Status(fiber.StatusAccepted).SendString("Please pay the fee first")
	} else {
		return c.Status(fiber.StatusForbidden).SendString("Active connection is not possible")
	}
}
