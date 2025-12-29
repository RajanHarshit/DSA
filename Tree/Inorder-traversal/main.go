package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversalWithStack(root *TreeNode) []int {
	result := []int{}
	stack := []*TreeNode{}
	curr := root

	for curr != nil || len(stack) > 0 {
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}

		curr = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, curr.Val)

		curr = curr.Right
	}

	return result
}

func inorderTraversalWithRecursion(root *TreeNode) []int {
	result := []int{}
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		result = append(result, node.Val)
		inorder(node.Right)
	}
	inorder(root)
	return result
}

func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 5}
	root.Left.Left = &TreeNode{Val: 8}
	root.Left.Right = &TreeNode{Val: 9}
	root.Right = &TreeNode{Val: 2}
	root.Right.Left = &TreeNode{Val: 3}

	fmt.Println(inorderTraversalWithStack(root))
	fmt.Println(inorderTraversalWithRecursion(root))
}
