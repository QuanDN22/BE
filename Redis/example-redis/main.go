package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type User struct {
	Name string `json:"name"`
	Age  int `json:"age"`
}

var ctx = context.Background()
var redisClient *redis.Client

func connectRedis(ctx context.Context) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// Ping Redis to check if the connection is working
	pong, err := client.Ping(ctx).Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(pong)

	redisClient = client
}

func setToRedis(ctx context.Context, key, val string) {
	err := redisClient.Set(ctx, key, val, 0).Err()
	if err != nil {
		fmt.Println(err)
	}
}

func getFromRedis(ctx context.Context, key string) string {
	val, err := redisClient.Get(ctx, key).Result()
	if err != nil {
		fmt.Println(err)
	}

	return val
}

func getAllKeys(ctx context.Context, key string) []string {
	keys := []string{}

	iter := redisClient.Scan(ctx, 0, key, 0).Iterator()
	for iter.Next(ctx) {
		keys = append(keys, iter.Val())
	}
	if err := iter.Err(); err != nil {
		panic(err)
	}

	return keys
}

func main() {
	// connect redis
	connectRedis(ctx)

	setToRedis(ctx, "name", "redis-test")
	setToRedis(ctx, "name2", "redis-test-2")

	val := getFromRedis(ctx, "name")
	fmt.Printf("First value with name key : %s \n", val)

	values := getAllKeys(ctx, "name*")
	fmt.Printf("All values : %v \n", values)

	values = getAllKeys(ctx, "")
	fmt.Printf("All values : %v \n", values)

	session := map[string]string{"name": "John", "surname": "Smith", "company": "Redis", "age": "29"}
	for k, v := range session {
		err := redisClient.HSet(ctx, "user-session:123", k, v).Err()
		if err != nil {
			panic(err)
		}
	}

	userSession, err := redisClient.HGetAll(ctx, "user-session:123").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(userSession)

	//
	user, err := json.Marshal(User{Name: "Quan", Age: 10})
	if err != nil {
		panic(err)
	}
	err = redisClient.Set(ctx, "id1234", user, 0).Err()
	if err != nil {
		panic(err)
	}
	val, err = redisClient.Get(ctx, "id1234").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(val)
}
