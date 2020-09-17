package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

func main() {
	opt, err := redis.ParseURL("redis://192.168.3.115:6379/0")
	if err != nil {
		panic(err)
	}
	rdb := redis.NewClient(opt)

	ud := map[string]interface{}{
		"Username":        "test",
		"LoginErrorTimes": 0,
		"Status":          "lock",
	}
	res := rdb.Set("userRd", ud, 30*time.Minute)
	fmt.Println(res.Val())

}
