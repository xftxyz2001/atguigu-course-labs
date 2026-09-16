package main

import (
	"fmt"
)

func main() {
	m1 := make(map[int]int)
	m1[10] = 100
	m1[100] = 10
	m1[8000] = 88
	m1[40] = 44
	m1[99] = 66
	// fmt.Println(m1)

	for k, v := range m1 {
		fmt.Println(k, ":", v)
	}

}
