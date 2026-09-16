package main

import (
	"fmt"
	"time"
)

// 计时器
func timer(f func()) {
	start := time.Now()
	f()
	end := time.Now()
	fmt.Printf("耗时：%v\n", end.Sub(start))
}

func test() {
	sum := 0
	for i := 0; i < 100000000; i++ {
		sum += i
	}
	// fmt.Println(sum)
}

func main() {
	now := time.Now()
	fmt.Println(now.Format("2006年1月2日 15时4分5秒"))

	// time.Duration
	// a := time.Second / 10
	// fmt.Printf("%T\n", a)

	// i := 0
	// for {
	// 	i++
	// 	fmt.Println(i)
	// 	// time.Sleep(time.Second * 0.1)
	// 	time.Sleep(time.Second / 10)
	// }

	// 将时间戳转换为时间格式
	var maxint64 int64 = 1<<63 - 1
	u := time.Unix(0, maxint64)
	fmt.Println(u)

	timer(test)

}
