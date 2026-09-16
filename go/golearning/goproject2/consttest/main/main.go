package main

import "fmt"

func main() {
	// const (
	// 	a = 1
	// 	b = iota
	// 	c = iota
	// 	d
	// )
	const (
		a = iota
		b
		c, d = iota, iota
	)
	fmt.Println(a, b, c, d)

}
