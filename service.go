package main

import (
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis"
)

var client = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "",
	DB:       0,
})

func GetAll() ([]byte, error) {
	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	return json.Marshal(Counter{})
}

func Get(id string) []byte {
	return []byte("{}")
}

func Increment(id string) []byte {
	return []byte("{}")
}

type Counter struct {
	value int
}
