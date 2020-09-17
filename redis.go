package main

import (
	"github.com/go-redis/redis"
	"log"
	"net"
	"time"
	"whiskey/conf"
)

var rc *redis.Client

func init() {
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
			return netDialer.Dial(conf.RedisConf["type"], conf.RedisConf["address"])
		},
	}

	//go-redis项目中的NewClient 函数自动创建 连接池，states结构体中可以查看连接池的状态信息
	rc := redis.NewClient(&redisOption)
	pong, err := rc.Ping().Result()
	if err != redis.Nil {
		log.Fatal("redis no found!")
	} else {
		log.Fatal("pong: ", pong)
	}
}
