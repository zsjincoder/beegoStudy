package database

import (
	"fmt"
	"github.com/go-redis/redis"
)

var Client *redis.Client

func RegisterRedis() {
	Client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",  // no password set
		DB:       0,   // use default DB
	})
	pong , err :=Client.Ping().Result()
	fmt.Println(pong,err)
}
