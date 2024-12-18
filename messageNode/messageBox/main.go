package main

import (
	"SensorManager/messageNode/data"
	"SensorManager/messageNode/database"
	"SensorManager/messageNode/rabbitMQ"
	"SensorManager/messageNode/utils"
	"encoding/json"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"time"
)

func main() {
	conn, err := amqp.DialConfig("amqp://guest:guest@localhost:5672/", amqp.Config{
		Heartbeat: 10 * time.Minute,
	})
	if err != nil {
		//TODO log the error
	}
	defer func(conn *amqp.Connection) {
		err := conn.Close()
		if err != nil {
			//TODO log the error
		}
	}(conn)
	ch, err := conn.Channel()
	if err != nil {
		//TODO log the error
	}
	defer func(ch *amqp.Channel) {
		err := ch.Close()
		if err != nil {

		}
	}(ch)
	// Declare a queue
	q, err := ch.QueueDeclare(
		rabbitMQ.Message, // queue name
		true,             // durable
		false,            // delete when unused
		false,            // exclusive
		false,            // no-wait
		nil,              // arguments
	)
	// Set up a consumer
	messages, err := ch.Consume(
		q.Name, // queue name
		"",     // consumer tag
		false,  // auto-acknowledge
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // arguments
	)
	fmt.Printf("Listening to queue '%s'...\n", rabbitMQ.Message)

	// Listen for messages indefinitely
	for msg := range messages {
		if msg.MessageId == rabbitMQ.SendMessage {
			db := utils.ConnectDatabase()
			user_wallet := db.Read(database.ACTIVE_USER)
			var message data.Message
			err = json.Unmarshal(msg.Body, &message)
			print(user_wallet)
			//TODO Send message to private node
		} else if msg.MessageId == rabbitMQ.EndChat {
		}
	}
}
