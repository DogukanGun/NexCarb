package database

import (
	"SensorManager/database"
	"SensorManager/test/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReadingNotCreatedObject(t *testing.T) {
	//Arrange
	db := database.Database{}
	key := utils.GenerateRandomString(8)
	db.Connect("localhost:6379", "")

	//Act
	value := db.Read(key)

	//Assert
	assert.Equal(t, value, "")
}
