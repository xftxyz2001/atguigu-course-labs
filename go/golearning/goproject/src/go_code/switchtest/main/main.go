package main

import "fmt"

func main() {

	switch ch := 'a'; ch {
	case 'a':
		fmt.Println("星期一")
		// break //go语言中switch默认有break，如果不想要break，可以使用fallthrough
	case 'b':
		fmt.Println("星期二")
	case 'c':
		fmt.Println("星期三")
	case 'd':
		fmt.Println("星期四")
	case 'e':
		fmt.Println("星期五")
	case 'f':
		fmt.Println("星期六")
	case 'g':
		fmt.Println("星期日")
	default:
		fmt.Println("输入有误")

	}

	var x interface{}
	var y = 10.0
	x = y

	switch i := x.(type) {
	case nil:
		fmt.Println(i)
	case float64:
		fmt.Println("float64", i)

	}
}
