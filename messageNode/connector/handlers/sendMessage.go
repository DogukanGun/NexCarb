package connectorHandlers

import (
	"SensorManager/messageNode/database"
	rabbitMQ2 "SensorManager/messageNode/rabbitMQ"
	"crypto/ed25519"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2/log"
	amqp "github.com/rabbitmq/amqp091-go"
)

type SendMessageRequest struct {
	Receiver  string `json:"receiver"`
	Message   string `json:"message"`
	Signature string `json:"signature"`
}

type Message struct {
	Receiver string `json:"receiver"`
	Message  string `json:"message"`
}

func SendMessageHandler(c *fiber.Ctx) error {
	p := new(SendMessageRequest)
	if err := c.BodyParser(p); err != nil {
		return err
	}
	db, ok := c.Locals("db").(*database.Database)
	if !ok {
		return c.Status(fiber.StatusInternalServerError).SendString("Database connection error")
	}
	publicKey := db.Read("active_user")
	sigBytes, _ := base64.StdEncoding.DecodeString(p.Signature)
	pubKeyBytes, _ := base64.StdEncoding.DecodeString(publicKey)
	signatureVerification := ed25519.Verify(pubKeyBytes, []byte(p.Message), sigBytes)
	if !signatureVerification {
		return c.Status(fiber.StatusInternalServerError).SendString("Signature is not valid")
	}
	sendListenCommand(p.Receiver, p.Message)
	return c.Status(fiber.StatusForbidden).SendString("the message is sent.")
}

func sendListenCommand(address string, message string) {
	// Connect to RabbitMQ server
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")

	defer func(conn *amqp.Connection) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)

	// Create a channel
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer func(ch *amqp.Channel) {
		err := ch.Close()
		if err != nil {

		}
	}(ch)

	// Declare a queue (to ensure it exists)
	_, err = ch.QueueDeclare(
		rabbitMQ2.Message, // name of the queue
		true,              // durable
		false,             // delete when unused
		false,             // exclusive
		false,             // no-wait
		nil,               // arguments
	)
	msg := Message{
		Receiver: address,
		Message:  message,
	}

	// Serialize the struct to JSON
	body, err := json.Marshal(msg)
	if err != nil {
		log.Fatalf("Failed to serialize message: %v", err)
	}

	// Publish a message to the queue
	err = ch.Publish(
		"",                // exchange (empty string for the default exchange)
		rabbitMQ2.Message, // routing key (queue name)
		false,             // mandatory
		false,             // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			MessageId:   rabbitMQ2.SendMessage,
			Body:        body,
		},
	)
	fmt.Printf("Sent message: %s\n", address)
}
