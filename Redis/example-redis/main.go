package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// Ping Redis to check if the connection is working
	v, err := rdb.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(v)

	// err := rdb.Set(ctx, "bike:1", "Process 134", 0).Err()
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("OK")

	// value, err := rdb.Get(ctx, "bike:1").Result()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("The name of the bike is %s", value)

	// err = rdb.Append(ctx, "quan", "dinh00").Err()
	// if err != nil {
	// 	panic(err)
	// }

}
