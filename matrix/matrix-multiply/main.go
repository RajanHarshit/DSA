package main

import (
	"fmt"
)

func multiMatrix(a, b [][]int) [][]int {
	m := len(a)
	n := len(b[0])

	if len(b) != m {
		fmt.Println("Multiplication is not possible")
	}

	// Initialise result
	result := make([][]int, m)
	for i := range result {
		result[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < len(a[i]); k++ {
				result[i][j] += a[i][k] * b[k][j]
			}

		}
	}
	return result
}
func main() {
	a := [][]int{
		{1, 2, 5},
		{3, 4, 6},
		{7, 8, 9},
	}
	b := [][]int{
		{10, 20, 50},
		{30, 40, 60},
		{70, 80, 90},
	}

	out := multiMatrix(a, b)
	for _, val := range out {
		fmt.Printf("%d\t", val)
		fmt.Println()
	}
}
