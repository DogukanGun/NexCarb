package database

import "github.com/redis/go-redis/v9"

func (db *Database) Connect(connectionString string, password string) {
	db.Rdb = redis.NewClient(&redis.Options{
		Addr:     connectionString,
		Password: password, // no password set
		DB:       0,        // use default DB
	})
}
