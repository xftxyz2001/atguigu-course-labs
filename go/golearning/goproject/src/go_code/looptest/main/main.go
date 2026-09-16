package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// for i := 0; i < 10; i++ {
	// 	fmt.Print(i)
	// }
	// fmt.Println()

	// j := 1
	// for j < 10 {
	// 	fmt.Print(j)
	// 	j++
	// }
	// fmt.Println()

	// k := 0
	// for {
	// 	if k > 9 {
	// 		break
	// 	}
	// 	fmt.Print(k)
	// 	k++
	// }

	// var str string = "hello world 你好，世界"
	// for i := 0; i < len(str); i++ {
	// 	fmt.Printf("i = %d, v = %c\n", i, str[i])
	// }
	// for i, v := range str {
	// 	fmt.Printf("i = %d, v = %c\n", i, v)
	// }

	totalLevel := 10

	for i := 1; i <= totalLevel; i++ {
		//在打印*前先打印空格
		for k := 1; k <= totalLevel-i; k++ {
			fmt.Print(" ")
		}

		//j 表示每层打印多少*
		for j := 1; j <= 2*i-1; j++ {
			if j == 1 || j == 2*i-1 || i == totalLevel {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}

		}
		fmt.Println()
	}

	var count = 0
	// rand.Seed(time.Now().Unix())
	for {
		count++
		// 生成一个随机数
		if rand.Intn(100)+1 == 99 {
			break
		}

	}
	fmt.Println(count)

// l:
// 	for i := 0; i < 10; i++ {
// 		for j := 0; j < 10; j++ {
// 			fmt.Println(j)
// 			if j > 3 {
// 				break l
// 			}
// 		}
// 	}

}
