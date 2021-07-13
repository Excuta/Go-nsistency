package repo

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/go-redis/redis"
	"github.com/jackc/pgx/v4"
)

var client = redis.NewClient(&redis.Options{
	Addr:     "redis:6379",
	Password: "",
	DB:       0,
})

const counterKey = "counter"
const expirySeconds = 30 * time.Second

var pgClient, pgerr = pgx.Connect(context.Background(), "postgres://yahia:2472BvZFgUNrof@postgres:5432/counter_db")

func GetCounter() (int64, error) {
	value, redisError := client.Get(counterKey).Result()
	fmt.Printf("Cached Value %q\n", value)
	res, parseError := strconv.ParseInt(value, 10, 64)
	if redisError != nil || parseError != nil {
		res = getFreshCounter()
		fmt.Printf("Fresh counter %v\n", res)
		client.Set(counterKey, res, expirySeconds)
	}
	return res, nil
}

func Increment() error {
	if pgerr != nil {
		return pgerr
	}
	go pgClient.QueryRow(context.Background(), "update counters set value = value +1").Scan()
	return nil
}

// assumes a value will always return from db
func getFreshCounter() int64 {
	var value int64
	pgClient.QueryRow(context.Background(), "select value from counters limit 1").Scan(&value)
	return value
}
