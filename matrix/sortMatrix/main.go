package main

import (
	"fmt"
	"sort"
)

func sortedMatrix(matrix [][]int) [][]int {
	n := len(matrix)

	var flat []int
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			flat = append(flat, matrix[i][j])
		}
	}

	sort.Ints(flat)

	idx := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			matrix[i][j] = flat[idx]
			idx++
		}
	}
	return matrix
}

func main() {

	matrix := [][]int{
		{10, 2},
		{3, 6},
	}

	fmt.Println(sortedMatrix(matrix))

}
