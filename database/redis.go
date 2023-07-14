package database

import (
	"fmt"
	"go-curd-redis-example/config"

	"github.com/go-redis/redis/v8"
)

func ConnectionRedisDb(config *config.Config) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr: config.RedisUrl,
	})

	fmt.Println("Connect db redis success!!!!")

	return rdb
}
