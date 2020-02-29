package rmodels

import (
	"log"

	"github.com/go-redis/redis"
)

// Redisdb ...
var Redisdb *redis.Client

func init() {
	Redisdb = redis.NewClient(&redis.Options{
		Addr:     "192.168.204.222:6379",
		Password: "",
		DB:       0,
	})

	_, err := Redisdb.Ping().Result()
	if err != nil {
		log.Printf("ping redis error, Failed code: %#v", err)
	}
}
