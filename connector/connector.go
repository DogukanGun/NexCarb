package main

import (
	"SensorManager/connector/handlers"
	connectormiddleware "SensorManager/connector/middleware"
	connectorUtils "SensorManager/connector/utils"
	"SensorManager/utils"
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

	utils.RunWithHandlingError(app.Listen("localhost:3000"))

}
