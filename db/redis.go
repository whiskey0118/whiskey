package db

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v7"
)

type Store interface {
	Set()
	Get()
}

type CTest struct {
	Name string
	Age  int
}

type Ha struct {
	St   Store
	Name string
}

func (c *CTest) Set() {

}
func (c *CTest) Get() {

}

func test() {
	//c :=
}

var ctx = context.Background()

func ExampleClient() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	err := rdb.Set("key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get("key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := rdb.Get("key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
	// Output: key value
}
