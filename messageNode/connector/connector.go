package main

import (
	"SensorManager/messageNode/connector/utils"
	"SensorManager/utils"
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
	app.Post("/payment", connectorHandlers.connectorHandlers.VerifyPaymentHandler)
	app.Post("/send/message", connectorHandlers.SendMessageHandler)
	app.Post("/user/active", connectorHandlers.IsActiveUserHandler)

	//Run app
	utils.RunWithHandlingError(app.Listen("localhost:3000"))
}
