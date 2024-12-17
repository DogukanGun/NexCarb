package main

import (
	connectorHandlers "SensorManager/messageNode/connector/handlers"
	connectormiddleware "SensorManager/messageNode/connector/middleware"
	"SensorManager/messageNode/utils"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName:           "connector",
		EnablePrintRoutes: true,
	})

	//Connect database
	//Connector endpoints must use this instance
	db := utils.ConnectDatabase()
	app.Use(connectormiddleware.DatabaseMiddleware(&db))
	app.Post("/connect", connectorHandlers.ConnectRequestHandler)
	app.Post("/payment", connectorHandlers.VerifyPaymentHandler)
	app.Post("/message/send", connectorHandlers.SendMessageHandler)
	app.Post("/user/active", connectorHandlers.IsActiveUserHandler)
	app.Post("/message/receive", connectorHandlers.ReceiveMessageHandler)

	//Run app
	utils.RunWithHandlingError(app.Listen("localhost:3000"))
}
