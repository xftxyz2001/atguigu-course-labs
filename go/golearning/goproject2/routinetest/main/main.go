package main

import (
	"fmt"
	"strconv"
	"time"
)

func test() {
	for i := 0; i <= 10; i++ {
		fmt.Println("Hello World" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}

func main() {
	go test()
	for i := 0; i <= 10; i++ {
		fmt.Println("Hello Golang" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}
