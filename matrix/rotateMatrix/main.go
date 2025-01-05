package main

import "fmt"

func rotateMatrixClockwise(matrix [][]int) {

	n := len(matrix)
	previous := matrix[0][0]

	// First row rotation to right
	for i := 0; i < n-1; i++ {
		matrix[0][i], previous = previous, matrix[0][i]
	}

	// Last column moving down
	for i := 0; i < n-1; i++ {
		matrix[i][n-1], previous = previous, matrix[i][n-1]
	}

	// Move Bottom row to left
	for i := n - 1; i > 0; i-- {
		matrix[n-1][i], previous = previous, matrix[n-1][i]
	}

	// Move first column
	for i := n - 1; i > 0; i-- {
		matrix[i][0], previous = previous, matrix[i][0]
	}
	matrix[0][0] = previous

}

func PrintMatrix(matrix [][]int) {
	for _, row := range matrix {
		for _, element := range row {
			fmt.Printf("%d\t", element)
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
	PrintMatrix(matrix)
	rotateMatrixClockwise(matrix)

	fmt.Println("\nAfter rotation of matrix::")
	PrintMatrix(matrix)

}
