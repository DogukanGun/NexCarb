package connectorUtils

import (
	"SensorManager/database"
	"SensorManager/env"
	"os"
)

func ConnectDatabase() database.Database {
	databaseUri := os.Getenv(env.DATABASE_URI)
	databasePassword := os.Getenv(env.DATABASE_PASSWORD)
	db := database.Database{}
	db.Connect(databaseUri, databasePassword)
	return db
}
