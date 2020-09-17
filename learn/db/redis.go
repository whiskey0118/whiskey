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

	rdb.HSet("ha", "fu", "saa")
	fmt.Println(rdb.HGet("ha", "fu"))

}

func main() {
	ExampleMewClient()
}
