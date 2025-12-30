package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func leavesNode(root *TreeNode) {
	if root == nil {
		return
	}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node.Left == nil && node.Right == nil {
			fmt.Print(node.Val, " ")
		}

		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}

	}
	fmt.Printf("\n")
}

func getLeaves(root *TreeNode) []int {
	leaves := []int{}

	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		if node.Left == nil && node.Right == nil {
			leaves = append(leaves, node.Val)
			return
		}
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	return leaves
}

func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}
	root.Right.Right = &TreeNode{Val: 6}

	leavesNode(root)             // BFS
	fmt.Println(getLeaves(root)) // Recursive
}
