package main

import "fmt"

var Age int = 50

// Name:="123"
var Name string = "jack~"

// var (
// 	age  int    = 50
// 	name string = "jack~"
// )

func main() {
	{
		a := 1
		fmt.Println(a)
	}
	// fmt.Println(a)

	i := 1
	for i := 0; i < 1; i++ {
		i := 2
		fmt.Println(i)
	}
	fmt.Println(i)

}
