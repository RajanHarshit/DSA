package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxHeightRecursive(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := maxHeightRecursive(root.Left)
	right := maxHeightRecursive(root.Right)

	if left > right {
		return left + 1
	}
	return right + 1
}

func maxHeightBFSQueue(root *TreeNode) int {
	if root == nil {
		return 0
	}
	depth := 0
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		levelsize := len(queue)
		depth++
		for i := 0; i < levelsize; i++ {
			node := queue[0]
			queue = queue[1:]

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	return depth
}

func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}
	root.Right.Right = &TreeNode{Val: 6}

	fmt.Println(maxHeightRecursive(root)) // recursive DFS
	fmt.Println(maxHeightBFSQueue(root))  // BFS queue
}
