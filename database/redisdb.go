package database

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
	"github.com/meloalright/guora/configuration"
)

var RDB *redis.Client
var ctx = context.Background()

func init() {

	RDB = redis.NewClient(&redis.Options{
		Addr:     configuration.C.Redis.Addr,
		Password: configuration.C.Redis.Password,
		DB:       configuration.C.Redis.Db,
	})

	if dbsize, err := RDB.DBSize(ctx).Result(); err != nil {
		log.Println("[redis]: error", err)
		panic("failed to connect redis")
	} else {
		log.Println("[redis]: dbsize", dbsize)

	}

}
