package database

import (
	"fmt"

	"github.com/go-redis/redis"
)

var RedisClient *redis.Client

func CreateRedisClient() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	pong, err := RedisClient.Ping().Result()
	fmt.Println(pong, err)
}
