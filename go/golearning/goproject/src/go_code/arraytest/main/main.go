package main

import "fmt"

func main() {
	// n := 12
	// var arr [n]int
	// fmt.Println(arr)

	// []int
	// 0xc000008048
	// 0xc000010138
	var arr []int = []int{1, 2, 3}
	fmt.Printf("%T\n", arr)
	fmt.Printf("%p\n%p\n", &arr, &arr[0])
	fmt.Println(arr)

	var arr2 [3]int = [3]int{1, 2, 3}
	fmt.Printf("%T\n", arr2)
	fmt.Printf("%p\n%p\n", &arr2, &arr2[0])
	fmt.Println(arr2)
	// [3]int
	// 0xc000010150
	// 0xc000010150

	var names = [...]string{2: "tom", "jack", 1: "mary"}
	fmt.Println(names)

	for index, value := range names {
		fmt.Printf("index=%v value=%v\n", index, value)
	}

}
