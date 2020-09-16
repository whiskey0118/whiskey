package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis"
)

var ctx = context.Background()

func ExampleMewClient() {
	opt, err := redis.ParseURL("redis://192.168.3.115:6379/0")
	if err != nil {
		panic(err)
	}
	rdb := redis.NewClient(opt)
	val, err := rdb.Get("username").Result()
	if err != nil {
		if err == redis.Nil {
			fmt.Println("key does not exist")
			return
		}
		panic(err)
	}
	fmt.Println(val)

}

func main() {
	ExampleMewClient()
}
