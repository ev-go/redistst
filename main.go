package main

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	userid := "user#125"
	currentusertoken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBVFRFTlRJT04hIjoi0J_RgNC40LLQtdGCLCDQnNCw0LrRgSA6KSIsIkRhdGEgYW5zd2VyIGlzIjoiMjExIiwiVG9rZW4gcmVxdWVzdCBhdCI6IjIwMjItMDUtMTJUMjI6MDI6MDMuNDIzNTc1NCswNTowMCIsImFkbWluIHBlcm1pc3Npb25zPyI6Im1heWJlIiwiZXhwIjoxNjUyMzc1NTIzLCJsb2dpbiI6InJvb3QifQ.9do8soXtimGxr9TDAd6EI2W0l-95U0SSJD_5GPz4kMA"

	node := rdb.Set(ctx, userid, currentusertoken, 0).Err()
	if node != nil {
		panic(node)
	}

	// err = rdb.Set(ctx, "key2", "74", 0).Err()
	// if err != nil {
	// 	panic(err)
	// }

	val, node := rdb.Get(ctx, userid).Result()
	if node == redis.Nil {
		fmt.Println("key1 does not exist")
	} else if node != nil {
		panic(node)
	} else {
		fmt.Println(userid, val)
	}
	// val, err := rdb.Get(ctx, "key").Result()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("key", val)

	val2, node := rdb.Get(ctx, "key2").Result()
	if node == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if node != nil {
		panic(node)
	} else {
		fmt.Println("key2", val2)
	}
	// Output: key value
	// key2 does not exist

}
