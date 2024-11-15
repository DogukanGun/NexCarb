package database

import "context"

func (db *Database) Write(key string, value string) (err error) {
	ctx := context.Background()
	err = db.Rdb.Set(ctx, key, value, 0).Err()
	if err != nil {
		//TODO Error logging
		panic(err)
	}
	defer ctx.Done()
	return
}
