package database

import (
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
)

func (db *Database) Read(key string) string {
	ctx := context.Background()
	val, err := db.Rdb.Get(ctx, key).Result()
	if err != nil {
		//TODO Error logging
		// Check if the error is due to the key not existing
		if errors.Is(err, redis.Nil) { // redis.Nil is the error returned when a key doesn't exist
			return ""
		}
		// Return any other error to the caller
		panic(err)
	}
	defer ctx.Done()
	return val
}
