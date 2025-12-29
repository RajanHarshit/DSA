package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func searchNode(root *TreeNode, target int) bool {
	if root == nil {
		return false
	}

	if root.Val == target {
		return true
	}
	return searchNode(root.Left, target) || searchNode(root.Right, target)
}

func searchNodeUsingBFSQueue(root *TreeNode, target int) bool {
	if root == nil {
		return false
	}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		node := queue[0]
		queue := queue[1:]

		if node.Val == target {
			return true
		}

		if node.Left != nil {
			queue = append(queue, node.Left)
		}

		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}

	return false
}

func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}
	root.Right.Right = &TreeNode{Val: 6}

	target := 4
	fmt.Println(searchNode(root, target))
}
