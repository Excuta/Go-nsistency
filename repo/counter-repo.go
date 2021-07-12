package repo

import (
	"context"
	"fmt"
	"strconv"

	"github.com/go-redis/redis"
	"github.com/jackc/pgx/v4"
)

var client = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "",
	DB:       0,
})

const expirySeconds = 30

var pgClient, pgerr = pgx.Connect(context.Background(), "postgres://yahia:3d%40k%24>JL%24V@192.168.25.43:5432/counter_db")

func GetCounter() (int, error) {
	value, err := client.Get("counter").Result()
	if err != nil {
		value = getFreshCounter()
	}
	return strconv.Atoi(value)
}

func Increment() error {
	if pgerr != nil {
		return pgerr
	}
	var greeting string
	pgClient.QueryRow(context.Background(), "select 'Hello, world!'").Scan(&greeting)
	fmt.Println(greeting)
	return nil
}

func getFreshCounter() string {
	return ""
}
