package main

import "fmt"

func main() {
	var arr = [...]int{3, 5, 8, 1, 9, 10, 4, 5}
	// 冒泡排序
	for i := 0; i < len(arr); i++ {
		hasSwap := false
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
				hasSwap = true
			}
		}
		if !hasSwap {
			break
		}
	}

	// 选择排序
	// for i := 0; i < len(arr); i++ {
	// 	minIndex := i
	// 	for j := i + 1; j < len(arr); j++ {
	// 		if arr[minIndex] > arr[j] {
	// 			minIndex = j
	// 		}
	// 	}
	// 	if minIndex != i {
	// 		arr[minIndex], arr[i] = arr[i], arr[minIndex]
	// 	}

	// }

	fmt.Println(arr)

	// 二分查找
	var arr2 = [...]int{1, 8, 10, 89, 1000, 1234}
	target := 1000
	left, right := 0, len(arr2)-1
	mid := (left + right) / 2
	for left <= right {
		if arr2[mid] == target {
			fmt.Printf("找到了，下标为%d\n", mid)
			break
		} else if arr2[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
		mid = (left + right) / 2
	}
	if left > right {
		fmt.Println("没有找到")
	}

}
