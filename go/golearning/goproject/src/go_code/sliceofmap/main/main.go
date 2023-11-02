package main

import "fmt"

func main() {
	som := make([]map[string]string, 2)
	m1 := make(map[string]string)
	m1["name"] = "zhangsan"
	m1["age"] = "18"
	som = append(som, m1)
	fmt.Println(som)
}
