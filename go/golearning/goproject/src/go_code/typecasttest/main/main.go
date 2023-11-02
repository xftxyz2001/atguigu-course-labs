package main

import (
	"fmt"
)

func main() {
	var _int int = 1
	var _int8 int8 = 2
	var _int64 int64 = 64
	// _int64+int64(_int8)
	i := int(_int8) + (_int) + int(_int64)
	fmt.Printf("%T", i)

	var i5 int = 5
	var i2 int = 2
	var i5div2 int = i5 / i2
	fmt.Println(i5div2)

	var f1 float32 = 5 / 2
	fmt.Println(f1)

	var f2 float32 = 5.0 / 2
	fmt.Println(f2)

	var f3 float32 = 5 / 2.0
	fmt.Println(f3)

	// var f4 float32 = float32(i5)/i2
	// fmt.Println(f4)

}
