package repo

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/go-redis/redis"
	"github.com/jackc/pgx/v4"
)

const BASE_URL = "192.168.1.108"

var redisClient = redis.NewClient(&redis.Options{
	Addr:     fmt.Sprintf("%v:6379", BASE_URL),
	Password: "",
	DB:       0,
})

var pgClient, pgerr = pgx.Connect(context.Background(), fmt.Sprintf("postgres://yahia:2472BvZFgUNrof@%v:5432/counter_db", BASE_URL))

const REDIS_COUNTER_KEY = "counter"
const REDIS_COUNTER_EXPIRY_MILLIS = 30 * time.Second

func init() {
	pgClient.Exec(context.Background(), "CREATE TABLE IF NOT EXISTS public.counters(id uuid NOT NULL, value bigint NOT NULL, CONSTRAINT \"unique id\" PRIMARY KEY (id), CONSTRAINT \"value non negative\" CHECK (value >= 0))")
	var count = 0
	pgClient.QueryRow(context.Background(), "SELECT COUNT(*) FROM public.counters;").Scan(&count)
	if count == 0 {
		pgClient.Exec(context.Background(), "INSERT INTO public.counters(id, value) VALUES (gen_random_uuid(), 0);")
	}
}

func GetCounter() (int64, error) {
	value, redisError := redisClient.Get(REDIS_COUNTER_KEY).Result()
	fmt.Printf("Cached Value %q\n", value)
	res, parseError := strconv.ParseInt(value, 10, 64)
	if redisError != nil || parseError != nil {
		res = getFreshCounter()
		fmt.Printf("Fresh counter %v\n", res)
		redisClient.Set(REDIS_COUNTER_KEY, res, REDIS_COUNTER_EXPIRY_MILLIS)
	}
	return res, nil
}

func Increment() error {
	if pgerr != nil {
		return pgerr
	}
	go pgClient.Exec(context.Background(), "update counters set value = value +1")
	return nil
}

func getFreshCounter() int64 {
	var value int64
	rows, err := pgClient.Query(context.Background(), "select value from counters limit 1")
	if err != nil {
		fmt.Println("Query failed: ", err)
	}
	if rows.Next() {
		rows.Scan(&value)
		fmt.Printf("Query successful: %v\n", value)
	} else {
		fmt.Println("Query failed: empty")
	}
	defer rows.Close()
	return value
}
