package main

import "fmt"

func test(arr [3]int) {
	arr[0] = 88
	// println(arr)
	// fmt.Println(arr)
}

func test2(arr *[3]int) {
	arr[0] = 99
	fmt.Println(arr)
}

func main() {
	var arr [3]int = [3]int{1, 2, 3}

	// var arr1 = arr
	// arr1[0] = 100
	// println(arr[0])
	// println(arr1[0])

	// test(arr)
	// println(arr[0])
	// println(arr1[0])

	test2(&arr)
	fmt.Println(arr)
}
