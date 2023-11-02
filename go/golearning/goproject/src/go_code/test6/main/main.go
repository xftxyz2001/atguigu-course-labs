package main

import "fmt"

// 打印金字塔
func printPyramid(totalLevel int, solid bool) {
	for i := 1; i <= totalLevel; i++ {
		//在打印*前先打印空格
		for k := 1; k <= totalLevel-i; k++ {
			fmt.Print(" ")
		}

		//j 表示每层打印多少*
		for j := 1; j <= 2*i-1; j++ {
			if solid {
				fmt.Print("*")
			} else if j == 1 || j == 2*i-1 || i == totalLevel {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}

		}
		fmt.Println()
	}
}

func main() {
	printPyramid(5, true)
	printPyramid(5, false)
}
