package main

import (
	"fmt"
	"math"
)

// 1)声明一个结构体 Circle,  字段为 radius
type Circle struct {
	Radius float64
}

// 2)声明一个方法 area 和 Circle 绑定，可以返回面积。
func (c Circle) area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) String() string {
	return fmt.Sprintf("Circle的半径是=%v, 圆的面积是=%v", c.Radius, c.area())
}

func main() {
	var c1 Circle = Circle{Radius: 1.0}
	fmt.Println(c1)
}
