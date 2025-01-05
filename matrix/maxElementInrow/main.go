package main

import "fmt"

func maxElemInRow(matrix [][]int) []int {
	var result []int
	n := len(matrix)

	for i := 0; i < n; i++ {
		largest := matrix[i][0]
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] > largest {
				largest = matrix[i][j]
			}
		}
		result = append(result, largest)
	}
	return result
}
func main() {
	matrix := [][]int{
		{1, 5, 3},
		{10, 1, 0},
	}

	out := maxElemInRow(matrix)
	fmt.Println(out)
}
