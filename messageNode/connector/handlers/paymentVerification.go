package connectorHandlers

import (
	"SensorManager/common/payment"
	"SensorManager/messageNode/database"
	"SensorManager/messageNode/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

type VerifyPayment struct {
	TxHash              string `json:"tx_hash"`
	SenderWallet        string `json:"sender_wallet"`
	IsPaymentWithSolana bool   `json:"chain"`
}

func VerifyPaymentHandler(c *fiber.Ctx) error {
	p := new(VerifyPayment)
	if err := c.BodyParser(p); err != nil {
		return err
	}
	//receive database instance
	db, ok := c.Locals("db").(*database.Database)
	if !ok {
		return c.Status(fiber.StatusInternalServerError).SendString("Database connection error")
	}
	if p.IsPaymentWithSolana {
		paymentVerified := payment.VerifyTheTransactions(p.TxHash, p.SenderWallet)
		if paymentVerified {
			utils.RunWithHandlingError(db.Write("user_"+p.SenderWallet+"_status", "approved"))
			utils.RunWithHandlingError(db.Write(database.ACTIVE_USER, p.SenderWallet))
			return c.Status(fiber.StatusAccepted).SendString("payment is verified")
		} else {
			return c.Status(fiber.StatusForbidden).SendString("payment is not verified")
		}
	} else {
		//TODO handle with ethena's sUSD
		return c.SendStatus(500)
	}
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
