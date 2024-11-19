package connectorHandlers

import (
	"context"
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

	//TODO deploy a smart contract for this device

	//Save instance to database
	conn, ok := c.Locals("db").(*pgx.Conn)
	if !ok {
		//TODO log error
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
		// TODO log the error
	}

	if commandTag.Insert() {
		return c.Status(fiber.StatusAccepted).SendString("registration is done")
	} else {
		return c.Status(fiber.StatusNotAcceptable).SendString("registration is failed")
	}
}
