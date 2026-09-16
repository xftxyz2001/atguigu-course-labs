package main

import (
	"fmt"
	"runtime"
)

func main() {
	i := runtime.NumCPU()
	fmt.Println(i)

	i2 := runtime.GOMAXPROCS(i - 1)
	fmt.Println(i2)

	// i3 := runtime.GOMAXPROCS(i - 1)
	// fmt.Println(i3)

}
