package main

import "fmt"

func main() {
	if a := false; a {
		fmt.Println(a)
		// fmt.Println(b)
	} else if b := false; b {
		fmt.Println(a)
		fmt.Println(b)
	} else {
		fmt.Println(a)
		fmt.Println(b)
	}
}
