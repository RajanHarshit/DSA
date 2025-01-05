package main

import "fmt"

func rotateMatrixInR_angle(matrix [][]int) [][]int {
	m := len(matrix)

	// Transpose of matrix
	for i := 0; i < m; i++ {
		for j := i + 1; j < m; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	// Reverse Matrix rows
	for i := 0; i < m; i++ {
		left, right := 0, m-1
		for left < right {
			matrix[i][left], matrix[i][right] = matrix[i][right], matrix[i][left]
			left++
			right--
		}
	}
	return matrix

}
func printMatrix(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Printf("%d\t", matrix[i][j])
		}
		fmt.Println()
	}
}
func main() {
	matrix := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}

	fmt.Println("Original Matrix::")
	printMatrix(matrix)

	fmt.Println("______________")

	printMatrix(rotateMatrixInR_angle(matrix))

}
