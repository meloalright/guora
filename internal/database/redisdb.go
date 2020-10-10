package database

import (
	"context"
	"log"

	"github.com/meloalright/guora/conf"

	"github.com/go-redis/redis/v8"
)

// RDB reference
var RDB *redis.Client
var ctx = context.Background()

func init() {

	RDB = redis.NewClient(&redis.Options{
		Addr:     conf.Config().Redis.Addr,
		Password: conf.Config().Redis.Password,
		DB:       conf.Config().Redis.Db,
	})

	if dbsize, err := RDB.DBSize(ctx).Result(); err != nil {
		log.Println("[redis]: error", err)
		panic("failed to connect redis")
	} else {
		log.Println("[redis]: dbsize", dbsize)

	}

}
