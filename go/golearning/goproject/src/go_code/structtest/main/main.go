package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	p := new(Person)
	p1 := *new(Person)
	fmt.Printf("%T %T", p, p1)

}
