package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversalInRecursion(root *TreeNode) []int {
	result := []int{}
	var preorder func(node *TreeNode)
	preorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		result = append(result, node.Val)
		preorder(node.Left)
		preorder(node.Right)
	}
	preorder(root)
	return result
}

func preorderTraversalInStack(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	result := []int{}
	stack := []*TreeNode{root}

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, node.Val)

		// first push right then left
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}
	return result

}

func preorderTraversalInDFS(root *TreeNode) []int {
	result := []int{}
	stack := []*TreeNode{}
	curr := root

	for curr != nil || len(stack) > 0 {

		for curr != nil {
			result = append(result, curr.Val)
			stack = append(stack, curr)
			curr = curr.Left
		}

		curr = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		curr = curr.Right
	}
	return result
}

func main() {
	root := &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 3}
	root.Right.Right = &TreeNode{Val: 6}
	root.Left = &TreeNode{Val: 2}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}

	fmt.Println(preorderTraversalInRecursion(root))
	fmt.Println(preorderTraversalInStack(root))
	fmt.Println(preorderTraversalInDFS(root))
}
