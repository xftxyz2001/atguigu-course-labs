package main

import "fmt"

func init() {
	fmt.Println(name)
	name = "tom"
	age = 10
	sayHi = func() {
		fmt.Println("hi~", name, age)
	}
}

var (
	name  string = "jack"
	age   int
	sayHi func()
)

func main() {
	sayHi()

}
