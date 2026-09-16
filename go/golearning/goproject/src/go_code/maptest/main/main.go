package main

import "fmt"

func main() {
	map1 := make(map[string]string)
	map1["name"] = "tom"
	map1["age"] = "18"
	fmt.Println(map1)

	s, b := map1["age"]
	if b {
		fmt.Println(s)
	}
}
