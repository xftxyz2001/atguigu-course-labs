package main

import (
	"fmt"
)

func main() {
	// 输入姓名和年龄
	var name string
	var age byte
	fmt.Print("请输入姓名：")
	fmt.Scanln(&name)
	fmt.Print("请输入年龄：")
	fmt.Scanln(&age)
	fmt.Printf("姓名：%v 年龄：%v\n", name, age)
}
