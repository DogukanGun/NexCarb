package connector

import (
	connectorHandlers "SensorManager/searchNode/connector/handlers"
	connectorMiddleware "SensorManager/searchNode/connector/middleware"
	"SensorManager/searchNode/connector/utils"
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
	"os"
)

func Run() {
	app := fiber.New(fiber.Config{
		AppName:           "receiver",
		EnablePrintRoutes: true,
	})

	//Create database instance and put in middleware
	conn, err := pgx.Connect(context.Background(), os.Getenv("POSTGRES_URI"))
	if err != nil {
		//TODO log errors
	}

	app.Use(connectorMiddleware.DatabaseMiddleware(conn))
	//Check database and table
	err = utils.CheckDatabaseSetup(conn)
	if err != nil {
		//TODO log the error
		return
	}

	//Define endpoints
	app.Post("/register", connectorHandlers.RegisterDevice)
	app.Post("/search", connectorHandlers.SearchHandler)
	app.Get("/check", connectorHandlers.CheckAssignmentHandler)
	app.Post("/answer", connectorHandlers.AnswerAssignments)

	//Run app
	err = app.Listen("localhost:3000")
	if err != nil {
		//TODO log the error
	}
	defer func(conn *pgx.Conn, ctx context.Context) {
		err := conn.Close(ctx)
		if err != nil {
			//TODO log errors
		}
	}(conn, context.Background())
}
