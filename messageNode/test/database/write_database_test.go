package database

import (
	"SensorManager/messageNode/database"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCanWriteDatabase(t *testing.T) {
	//Arrange
	db := database.Database{}
	db.Connect("localhost:6379", "")

	//Act
	err := db.Write("test", "test")

	//Assert
	assert.Equal(t, err, nil)

}
