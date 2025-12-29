package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversalInRecursion(root *TreeNode) []int {
	result := []int{}
	var postorder func(node *TreeNode)
	postorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		postorder(node.Left)
		postorder(node.Right)
		result = append(result, node.Val)
	}
	postorder(root)
	return result
}

func postorderTraversalInStack(root *TreeNode) []int {
	result := []int{}
	stack := []*TreeNode{}
	var lastvisited *TreeNode
	curr := root
	fmt.Println("curr__", curr)

	for curr != nil || len(stack) > 0 {
		if curr != nil {
			stack = append(stack, curr)
			fmt.Println("stack__", stack)
			curr = curr.Left
			fmt.Println("curr.Left___", curr)
		} else {
			peek := stack[len(stack)-1]
			fmt.Println("peek___", peek)
			if peek.Right != nil && lastvisited != peek.Right {
				curr = peek.Right
				fmt.Println("peek.Right___", curr)
			} else {
				result = append(result, peek.Val)
				fmt.Println("result___", result)
				lastvisited = peek
				fmt.Println("lastvisited___", lastvisited)
				stack = stack[:len(stack)-1]
				fmt.Println("stack__here_", stack)
			}
		}

		fmt.Println()
	}
	return result
}

func postorderTraversalInDFS(root *TreeNode) []int {
	result := []int{}
	return result
}

func main() {
	root := &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 3}
	root.Right.Right = &TreeNode{Val: 6}
	root.Left = &TreeNode{Val: 2}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}

	fmt.Println(postorderTraversalInRecursion(root))
	fmt.Println(postorderTraversalInStack(root))
	fmt.Println(postorderTraversalInDFS(root))
}
