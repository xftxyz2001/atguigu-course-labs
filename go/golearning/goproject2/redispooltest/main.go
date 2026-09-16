package main

import (
	"github.com/gomodule/redigo/redis"
)

var pool *redis.Pool

func init() {
	pool = &redis.Pool{
		MaxIdle:     8,   // 最大空闲链接数
		MaxActive:   0,   // 表示和数据库的最大链接数，0表示没有限制
		IdleTimeout: 100, // 最大空闲时间
		Dial: func() (redis.Conn, error) { // 初始化链接的代码，链接哪个ip的redis
			return redis.Dial("tcp", "localhost:6379")
		},
	}
}

func main() {
	conn := pool.Get()
	defer conn.Close()
	// c, err := redis.Dial("tcp", "localhost:6379")

}
