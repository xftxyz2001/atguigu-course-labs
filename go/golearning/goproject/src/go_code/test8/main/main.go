package main

import (
	_ "errors"
	"fmt"
)

func testError() {
	// err := errors.New("test")
	panic("123")
}

func test() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("err = ", err)
		}
		fmt.Println("finally")
	}()
	testError()
}

func main() {
	test()
	fmt.Println("main")
}
