package main

import "fmt"

func main() {
	var c1 chan int
	var c2 <-chan int
	var c3 chan<- int

	c1 = make(chan int, 10)

	c2 = c1
	c3 = c1

	// c1 = c2 //cannot use c2 (variable of type <-chan int) as chan int value in assignment
	// c3 = c2 //cannot use c2 (variable of type <-chan int) as chan<- int value in assignment

	// c1 = c3 //cannot use c3 (variable of type chan<- int) as chan int value in assignment
	// c2 = c3 //cannot use c3 (variable of type chan<- int) as <-chan int value in assignment

	fmt.Printf("c1=%v, &c1=%v\n", c1, &c1)
	fmt.Printf("c2=%v, &c2=%v\n", c2, &c2)
	fmt.Printf("c3=%v, &c3=%v\n", c3, &c3)
}
