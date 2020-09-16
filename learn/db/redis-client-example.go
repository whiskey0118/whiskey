package main

import (
	"github.com/go-redis/redis"
	"time"
)

func main() {
	opt, err := redis.ParseURL("redis://192.168.3.115:6379/0")
	if err != nil {
		panic(err)
	}
	rdb := redis.NewClient(opt)

	<-time.After(1 * time.Second)

}
