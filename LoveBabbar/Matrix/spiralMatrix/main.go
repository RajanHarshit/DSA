/*Solution - 1*/
package main

import "fmt"

// func SpiralMatrixTraverse(matrix [][]int) []int {
// 	result := []int{}
// 	if len(matrix) == 0 || len(matrix[0]) == 0 {
// 		return result
// 	}
// 	top, bottom, left, right := 0, len(matrix)-1, 0, len(matrix[0])-1
// 	for top <= bottom && left <= right {

// 		// Traverse from left to right
// 		for i := left; i <= right; i++ {
// 			result = append(result, matrix[top][i])
// 		}
// 		top++

// 		// Traverse from top to bottom
// 		for i := top; i <= bottom; i++ {
// 			result = append(result, matrix[i][right])
// 		}
// 		right--

// 		if top <= bottom {
// 			// Traverse from right to left
// 			for i := right; i >= left; i-- {
// 				result = append(result, matrix[bottom][i])
// 			}
// 			bottom--

// 		}

// 		if left <= right {
// 			// Traverse from bottom to top
// 			for i := bottom; i >= top; i-- {
// 				result = append(result, matrix[i][left])
// 			}
// 			left++
// 		}
// 	}
// 	return result
// }

func main() {
	matrix := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}

	//fmt.Println(SpiralMatrixTraverse(matrix))
	fmt.Println(SpiralMatrixTraverseAnotherWay(matrix))
}

func SpiralMatrixTraverseAnotherWay(matrix [][]int) []int {
	result := []int{}
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return result
	}

	top, bottom, left, right := 0, len(matrix)-1, 0, len(matrix[0])-1

	for top <= bottom && left <= right {

		// Traverse left to right
		for i := left; i <= right; i++ {
			result = append(result, matrix[top][i])
		}
		top++

		// Traverse from top to bottom
		for i := top; i <= bottom; i++ {
			result = append(result, matrix[i][right])
		}
		right--

		// Traverse from right to left
		for i := right; i >= left; i-- {
			result = append(result, matrix[bottom][i])
		}
		bottom--

		// Traverse from bottom to top
		for i := bottom; i >= top; i-- {
			result = append(result, matrix[i][left])
		}
		left++
	}
	return result
}
