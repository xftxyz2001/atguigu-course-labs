package main

import (
	// fmt
	"fmt"
	// redis
	"github.com/gomodule/redigo/redis"
)

func main() {
	// 链接
	c, err := redis.Dial("tcp", "localhost:6379")
	// 错误处理
	if err != nil {
		fmt.Println("conn redis failed,", err)
		return
	}
	// 关闭
	defer c.Close()
	// 执行
	_, err = c.Do("Set", "abc", 100)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 执行
	r, err := redis.Int(c.Do("Get", "abc"))
	if err != nil {
		fmt.Println("get abc failed,", err)
		return
	}
	// 打印
	fmt.Println(r)

}
