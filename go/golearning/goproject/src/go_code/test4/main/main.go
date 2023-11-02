package main

import "fmt"

func main() {
	fmt.Printf("%b", -2^2)

	if a, b := 1, 2; a == 1 && b == 2 {
		fmt.Println("true")
	}
	n1, n2 := 10, 5
	if sum := n1 + n2; sum%3 == 0 && sum%5 == 0 {
		fmt.Println("能被3和5整除")
	}
}
