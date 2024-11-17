package bootsrapHandlers

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
	"os"
)

type RegisterNewNodeRequest struct {
	NodeIP      string `json:"node_ip"`
	OwnerPubKey string `json:"owner_pub_key"`
	MessageFee  string `json:"message_fee"`
}

func RegisterNewNodeHandler(c *fiber.Ctx) error {
	p := new(RegisterNewNodeRequest)
	if err := c.BodyParser(p); err != nil {
		return err
	}

	//Connect Database
	conn, err := pgx.Connect(context.Background(), os.Getenv("POSTGRES_URI"))
	if err != nil {
		fmt.Printf("Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer func(conn *pgx.Conn, ctx context.Context) {
		err := conn.Close(ctx)
		if err != nil {
			//TODO log error
		}
	}(conn, context.Background())
	ctx := context.Background()
	args := pgx.NamedArgs{
		"nodeIP":      p.NodeIP,
		"ownerPubkey": p.OwnerPubKey,
		"messageFee":  p.MessageFee,
	}
	commandTag, err := conn.Exec(
		ctx,
		"INSERT INTO nodes (node_ip,owner_pubkey,message_fee) VALUES (@nodeIP,@ownerPubkey,@messageFee)",
		args,
	)
	defer ctx.Done()
	if commandTag.Insert() {
		return c.Status(fiber.StatusForbidden).SendString("the message is sent.")
	} else {
		return c.Status(fiber.StatusInternalServerError).SendString("Cannot be registered")

	}
}
