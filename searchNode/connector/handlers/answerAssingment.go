package connectorHandlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
	"log"
)

type AssignmentAnswer struct {
	UserContract string `json:"user_contract"`
	Answer       string `json:"answer"`
}

type AnswerAssignmentsRequest struct {
	Answers           []AssignmentAnswer `json:"answers"`
	NodeWalletAddress string             `json:"node_wallet_address"`
}

func AnswerAssignments(c *fiber.Ctx) error {
	p := new(AnswerAssignmentsRequest)
	if err := c.BodyParser(p); err != nil {
		return err
	}
	conn, ok := c.Locals("db").(*pgx.Conn)
	if !ok {
		return c.Status(fiber.StatusInternalServerError).SendString("database connection error")
	}
	for _, element := range p.Answers {
		//Update database status
		query := `
			UPDATE node_assignments
			SET is_active = 1
			WHERE node_wallet_address = $1 AND user_contract = $2;
		`
		// Execute the query
		_, err := conn.Exec(c.Context(), query, p.NodeWalletAddress, element.UserContract)
		if err != nil {
			log.Println("Error updating assignment:", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to update assignment",
			})
		}
		//TODO Send answer to user
	}

	return c.SendStatus(fiber.StatusAccepted)
}
