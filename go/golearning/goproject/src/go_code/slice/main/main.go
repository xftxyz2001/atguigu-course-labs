package main

import (
	"fmt"
)

func main() {
	// i := make([]int, 5)
	// fmt.Println(i)

	// var arr = [10]int{1, 2, 3, 4, 5, 6}
	// var slice = arr[2:3]
	// fmt.Println("cap(slice)", cap(slice))

	// fmt.Println("arr", arr)
	// fmt.Println("slice", slice)

	// slice = append(slice, 4, 5, 6)
	// // fmt.Println("cap(slice)", cap(slice))
	// fmt.Println("arr", arr)
	// fmt.Println("slice", slice)

	// src := arr[2:6] // 3,4,5,6
	// dst := arr[0:6] // 1,2,3,4,5,6

	// i := copy(dst, src)
	// fmt.Println(i)
	// fmt.Println(dst)
	// fmt.Println(src)
	// fmt.Println(arr) // [3 4 5 6 5 6 0 0 0 0]

	// src, dst = dst, src
	// i = copy(dst, src)
	// fmt.Println(i)
	// fmt.Println(dst)
	// fmt.Println(src)
	// fmt.Println(arr) //	[1 2 1 2 3 4 0 0 0 0]

	str := "hello world"
	// slice := str[1:]
	// fmt.Printf("%T\n", slice)
	// // slice[0] = 'a'
	// fmt.Println(slice)

	// 使用切片将hello改为你好，world改为世界
	slice := []rune(str)
	slice[0] = '你'
	slice[1] = '好'
	slice[6] = '世'
	slice[7] = '界'
	str = string(slice)
	fmt.Println(str) //你好llo 世界rld

}
