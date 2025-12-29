package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func parentNode(root *TreeNode, target int) *TreeNode {
	if root == nil {
		return nil
	}

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node.Left != nil {
			if node.Left.Val == target {
				return node
			}
			queue = append(queue, node.Left)
		}

		if node.Right != nil {
			if node.Right.Val == target {
				return node
			}
			queue = append(queue, node.Right)
		}
	}
	return nil
}

func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}
	root.Right.Right = &TreeNode{Val: 6}

	target := 4
	fmt.Println(parentNode(root, target).Val) // BFS
}
