package main

import "fmt"

type Point struct {
	x, y int
}

type Point3 struct {
	x, y, z int
}

func main() {
	var a interface{}
	var i int
	fmt.Scanln(&i)

	if i == 1 {
		a = Point{1, 2}
	} else {
		a = Point3{1, 2, 3}
	}

	// b:=Point(a)
	b := a.(Point)
	fmt.Println("b = ", b) // b = {1 2}

}
