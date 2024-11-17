package database

import (
	"SensorManager/messageNode/database"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReadDatabase(t *testing.T) {
	//Arrange
	db := database.Database{}
	key, value := "test", "test_correct"
	db.Connect("localhost:6379", "")
	err := db.Write(key, value)
	assert.Equal(t, err, nil)

	//Act
	valueFromDb := db.Read(key)

	//Assert
	assert.Equal(t, valueFromDb, value)
}
