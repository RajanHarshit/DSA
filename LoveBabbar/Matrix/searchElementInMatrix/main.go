package main

import "fmt"

func SearchElementInMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return true
	}
	m := len(matrix)
	n := len(matrix[0])
	low, high := 0, m*n-1

	for low <= high {
		mid := low + (high-low)/2
		midRow := mid / n
		midCol := mid % n
		midElement := matrix[midRow][midCol]
		fmt.Printf("midElement: %v\n", midElement)

		if midElement == target {
			return true
		}
		if midElement < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return false
}
func main() {
	matrix := [][]int{
		{1, 3, 5, 7},
		{10, 11, 12, 14},
		{17, 18, 23, 27},
	}
	target := 2
	if SearchElementInMatrix(matrix, target) {
		fmt.Println("Found")
	} else {
		fmt.Println("NotFound")
	}
}
