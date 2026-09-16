package main

import "fmt"

func main() {
	// var i1 int = 5
	// var i2 int = 2
	// fmt.Println(i1 % i2)
	// var i3 int
	// i3 += i1

	// day := 97
	var day int = 97
	week, day := day/7, day%7
	fmt.Println(week, day)

	var f float64 = 123.2
	// c := float64(5) / 9 * (f - 100)
	c := 5.0 / 9 * (f - 100)
	fmt.Println(c)

	var a int = 10
	var b float64 = 3.4
	b1 := float64(a) > b
	fmt.Println(b1)

	str := "hello"
	str += " world"
	fmt.Println(str)

	var (
		i1 = 1
		i2 = 2
	)
	i1, i2 = i2, i1
	fmt.Println(i1, i2)
	fmt.Println(5<<1, -5>>1)
	u := max(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println("u =", u)

}
