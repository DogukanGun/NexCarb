package database

import (
	"SensorManager/messageNode/database"
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCanConnectDatabase(t *testing.T) {
	//Arrange
	db := database.Database{}

	//Act
	db.Connect("localhost:6379", "")
	ctx := context.TODO()
	pong, err := db.Rdb.Ping(ctx).Result()

	//Assert
	assert.NotEmpty(t, pong)
	assert.Equal(t, err, nil)

}
