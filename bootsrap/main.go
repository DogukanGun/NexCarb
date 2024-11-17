package main

import (
	bootsrapHandlers "SensorManager/bootsrap/handlers"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName:           "bootstrap",
		EnablePrintRoutes: true,
	})
	app.Post("/register", bootsrapHandlers.RegisterNewNodeHandler)
	app.Get("/nodes", bootsrapHandlers.GetNodesHandler)
	err := app.Listen("localhost:3001")
	if err != nil {
		//TODO log error
	}
}
