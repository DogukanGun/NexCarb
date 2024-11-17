package bootsrapHandlers

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
	"os"
)

func GetNodesHandler(c *fiber.Ctx) error {
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
	rows, err := conn.Query(ctx, "SELECT * FROM nodes")
	if err != nil {
	}
	defer rows.Close()
	var rowSlice []RegisterNewNodeRequest
	for rows.Next() {
		var r RegisterNewNodeRequest
		err := rows.Scan(&r.NodeIP, &r.MessageFee, &r.OwnerPubKey)
		if err != nil {
		}
		rowSlice = append(rowSlice, r)
	}
	if err := rows.Err(); err != nil {
	}
	return c.JSON(rowSlice)
}
