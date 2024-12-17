package connectorHandlers

import (
	"SensorManager/common/logger"
	"context"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
)

type RegisterDeviceRequest struct {
	WalletAddress string `json:"wallet_address"`
	Country       string `json:"location"`
}

func RegisterDevice(c *fiber.Ctx) error {
	p := new(RegisterDeviceRequest)
	if err := c.BodyParser(p); err != nil {
		return err
	}

	//Save instance to database
	conn, ok := c.Locals("db").(*pgx.Conn)
	if !ok {
		logger.LogE("database object is missing")
		return errors.New("database object is missing")
	}

	ctx := context.Background()
	args := pgx.NamedArgs{
		"walletAddress": p.WalletAddress,
		"country":       p.Country,
		"smartContract": "",
	}
	commandTag, err := conn.Exec(
		ctx,
		"INSERT INTO search_node_registery (wallet_address,country,smart_contract) VALUES (@walletAddress,@country,@smartContract)",
		args,
	)
	defer ctx.Done()

	if err != nil {
		logger.LogE("Error while inserting data into database", err)
		return errors.New("database error")
	}

	if commandTag.Insert() {
		return c.Status(fiber.StatusAccepted).SendString("registration is done")
	} else {
		return c.Status(fiber.StatusNotAcceptable).SendString("registration is failed")
	}
}
