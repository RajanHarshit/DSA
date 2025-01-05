package main

import (
	"fmt"
)

func countAllSortedRows(matrix [][]int) int {
	count := 0
	for _, row := range matrix {
		if isSorted(row) {
			count++
		}
	}
	return count
}

func isSorted(row []int) bool {
	n := len(row)

	if n < 2 {
		return true
	}
	isIncreasing := true
	isDecreasing := true

	for i := 1; i < n; i++ {
		if row[i] <= row[i-1] {
			isIncreasing = false
		}
		if row[i] >= row[i-1] {
			isDecreasing = false
		}
	}

	return isIncreasing || isDecreasing

}
func main() {
	matrix := [][]int{
		{1, 2, 3, 4},
		{3, 2, 10, 4},
		{11, 9, 7, 5},
	}

	out := countAllSortedRows(matrix)
	fmt.Println(out)
}
