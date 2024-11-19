package connectorHandlers

import (
	"SensorManager/common/payment"
	"SensorManager/messageNode/env"
	"context"
	"database/sql"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
	"os"
)

type SearchRequest struct {
	SearchText        string `json:"search_text"`
	UserWalletAddress string `json:"user_wallet_address"`
	Country           string `json:"country"`
	TxHash            string `json:"tx_hash"`
}

func SearchHandler(c *fiber.Ctx) error {
	p := new(SearchRequest)
	if err := c.BodyParser(p); err != nil {
		return err
	}
	paymentVerified := payment.VerifyTheTransactions(p.TxHash, os.Getenv(env.PUBKEY))
	conn, ok := c.Locals("db").(*pgx.Conn)
	if !ok {
		return c.Status(fiber.StatusInternalServerError).SendString("database connection error ")

	}
	if !paymentVerified {
		return c.Status(fiber.StatusNotAcceptable).SendString("please make payment first")
	}
	// Select a random node from `search_node_registery` filtered by country if provided
	var nodeWalletAddress string
	query := "SELECT wallet_address FROM search_node_registery "
	var args []interface{}

	if p.Country != "" {
		query += "WHERE country = $1 "
		args = append(args, p.Country)
	}
	query += "ORDER BY RANDOM() LIMIT 1;"

	err := conn.QueryRow(context.Background(), query, args...).Scan(&nodeWalletAddress)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return c.Status(fiber.StatusNotFound).SendString("no nodes available")
		}
		return c.Status(fiber.StatusInternalServerError).SendString("failed to fetch node")
	}

	// Write the assignment to `node_assignments`
	insertQuery := `
		INSERT INTO node_assignments (node_wallet_address, search_text,user_contract,is_active)
		VALUES ($1, $2, $3, B'1');`
	_, err = conn.Exec(context.Background(), insertQuery, nodeWalletAddress, p.SearchText, p.UserWalletAddress)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("failed to assign node")
	}

	// Return the assigned node
	return c.JSON(fiber.Map{
		"assigned_node": nodeWalletAddress,
	})
}
