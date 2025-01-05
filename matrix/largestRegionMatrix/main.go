/*Largest region in matrix of 1s */
package main

import "fmt"

var Directions = []struct{ x, y int }{
	{-1, 0}, {1, 0}, {0, 1}, {0, -1},
	{1, 1}, {-1, 1}, {-1, -1}, {1, -1},
}

// function to perform DFS and count the size of region
func dfs(matrix [][]int, visited [][]bool, x, y int) int {
	// base case : out of bounds or already visited or not a 1s
	if x < 0 || x >= len(matrix) || y < 0 || y >= len(matrix[0]) || matrix[x][y] == 0 || visited[x][y] {
		return 0
	}

	// mark this cell as visited
	visited[x][y] = true

	//Start with current cell size
	size := 1

	// Explore all 8 direction
	for _, dir := range Directions {
		newX, newY := x+dir.x, y+dir.y
		size += dfs(matrix, visited, newX, newY)
	}
	return size
}

// function to finds largest region of 1s in the matrix
func largestRegion(matrix [][]int) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}

	// Create a visited matrix to track visited cell
	visited := make([][]bool, len(matrix))
	for i := range visited {
		visited[i] = make([]bool, len(matrix[0]))
	}

	largest := 0

	// Traverse entire matrix
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			// If the cell is 1 and has not been visited, start a DFS
			if matrix[i][j] == 1 && !visited[i][j] {
				regionSize := dfs(matrix, visited, i, j)
				if regionSize > largest {
					largest = regionSize
				}
			}
		}
	}
	return largest
}
func main() {
	// Example matrix with multiple regions of 1s
	matrix := [][]int{
		{1, 1, 0, 0, 0},
		{1, 1, 0, 1, 1},
		{0, 0, 0, 1, 1},
		{0, 1, 1, 1, 0},
		{1, 0, 0, 0, 1},
	}

	// Call the function to find the largest region
	result := largestRegion(matrix)

	// Print the result
	fmt.Println("The largest region size is:", result)
}
