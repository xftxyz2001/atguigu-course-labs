package main

import (
	"fmt"
	"reflect"
)

func test01(b interface{}) {
	t := reflect.TypeOf(b)
	fmt.Printf("t:%v\n", t)
	// reflect.Value.Kind()
	v := reflect.ValueOf(b)
	fmt.Printf("v:%v\n", v)

	a := v.Interface()
	fmt.Printf("a:%v\n", a)

	i := a.(int)
	fmt.Printf("i:%v\n", i)
}

func main() {
	test01(123)
	str:= "hello"
	fs := reflect.ValueOf(str)
	fs.SetString("world")
	fmt.Println("str=", str)
}
