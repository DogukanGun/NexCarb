package main

import (
	connectorHandlers "SensorManager/messageNode/connector/handlers"
	connectormiddleware "SensorManager/messageNode/connector/middleware"
	"SensorManager/messageNode/connector/utils"
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
	db := connectorUtils.ConnectDatabase()
	app.Use(connectormiddleware.DatabaseMiddleware(&db))
	app.Post("/connect", connectorHandlers.ConnectRequestHandler)
	app.Post("/payment", connectorHandlers.VerifyPaymentHandler)
	app.Post("/send/message", connectorHandlers.SendMessageHandler)
	app.Post("/user/active", connectorHandlers.IsActiveUserHandler)

	//Run app
	utils.RunWithHandlingError(app.Listen("localhost:3000"))
}
