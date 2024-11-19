package connectorHandlers

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
)

type CheckAssignmentResponse struct {
	UserWallet string `json:"user_wallet"`
	SearchText string `json:"search_text"`
}

func CheckAssignmentHandler(c *fiber.Ctx) error {
	walletAddress := c.Params("wallet_address")

	conn, ok := c.Locals("db").(*pgx.Conn)
	if !ok {
		//TODO log error
	}
	var response []CheckAssignmentResponse
	query := "SELECT user_wallet, search_texst FROM node_assignments WHERE node_wallet_address = $1 and is_active = 1"
	args := []interface{}{walletAddress}
	err := conn.QueryRow(context.Background(), query, args...).Scan(&response)
	if err != nil {
		//TODO log error
	}
	return c.JSON(response)
}
