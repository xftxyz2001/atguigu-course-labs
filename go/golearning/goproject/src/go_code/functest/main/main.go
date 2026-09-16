package main

import "fmt"

func test(a, b int) (int, int) {
	return a + b, a - b
}

// func cal(a int, b int) (sum int, sub int) {
// 	sum = a + b
// 	sub = a - b
// 	return
// }

func sum(a int, b int) int {
	return a + b
}

func sub(a int, b int) int {
	return a - b
}

type op func(int, int) int

func cal(f op, a int, b int) (res int) {
	fmt.Printf("f=%T\n", f)
	res = f(a, b)
	return
}

func main() {

	sm, sb := test(10, 20)
	fmt.Println("sum=", sm, "sub=", sb)

	a, b := 10, 20
	res := cal(sum, a, b)
	fmt.Println("res=", res)
	res = cal(sub, a, b)
	fmt.Println("res=", res)

	// println("sum=", sum, "sub=", sub)
	// sum2, sub2 := cal(1, 2)

}
