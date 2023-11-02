package main

import "fmt"

// 实现矩阵转置
func transpose(matrix [][]int) [][]int {
	for i := 0; i < len(matrix); i++ {
		for j := i; j < len(matrix[i]); j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	return matrix

}

// 打印矩阵
func printMatrix(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Printf("%v\t", matrix[i][j])
		}
		fmt.Println()
	}
}

func main() {
	// 测试
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	printMatrix(matrix)
	transpose(matrix)
	fmt.Println()
	printMatrix(matrix)

}
