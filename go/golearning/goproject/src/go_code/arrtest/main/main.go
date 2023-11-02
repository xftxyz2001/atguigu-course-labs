package main

import "fmt"

func test1() {
	// var arr [26]byte = [26]byte{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J',
	// 	'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T',
	// 	'U', 'V', 'W', 'X', 'Y', 'Z'}

	var arr [26]byte
	for i := 0; i < len(arr); i++ {
		arr[i] = 'A' + byte(i)
	}

	for i := 0; i < len(arr); i++ {
		fmt.Printf("%c ", arr[i])
	}

}

func mymax(arr [5]int) (index, value int) {
	index = 0
	for i := 0; i < 5; i++ {
		if arr[i] > arr[index] {
			index = i
			value = arr[i]
		}
	}
	return
}

func main() {
	test1()

	index, value := mymax([5]int{1, 2, 3, 4, 5})
	fmt.Println(index, value)
	// len:=1
	// l:=len([]int{1,2,3})
	// fmt.Println(l)
}
