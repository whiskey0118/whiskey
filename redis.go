package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"net"
	"time"
)

func main() {
	redisOption := redis.Options{
		Dialer: func() (conn net.Conn, err error) {
			netDialer := &net.Dialer{
				Timeout:       5 * time.Second,
				Deadline:      time.Time{},
				LocalAddr:     nil,
				DualStack:     false,
				FallbackDelay: 0,
				KeepAlive:     5 * time.Second,
				Resolver:      nil,
				Cancel:        nil,
				Control:       nil,
			}
			return netDialer.Dial("tcp", "127.0.0.1:6379")
		},

		OnConnect: func(conn *redis.Conn) error {
			fmt.Printf("conn=&s\n", conn)
			return nil
		},
	}

	//go-redis项目中的NewClient 函数自动创建 连接池，states结构体中可以查看连接池的状态信息
	rc := redis.NewClient(&redisOption)
	defer rc.Close()

	rc.Set("haha", "fuck", 10*time.Second)
	fmt.Println(rc.PoolStats())
}
