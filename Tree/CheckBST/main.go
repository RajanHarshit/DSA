package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func validate(node *TreeNode, min, max int64) bool {
	if node == nil {
		return true
	}
	if int64(node.Val) <= min || int64(node.Val) >= max {
		return false
	}
	return validate(node.Left, min, int64(node.Val)) && validate(node.Right, int64(node.Val), max)
}

func isBST(root *TreeNode) bool {
	return validate(root, math.MinInt64, math.MaxInt64)
}

func main() {
	/*
	   Example Tree (Valid BST)
	           10
	          /  \
	         5    15
	             /  \
	            11  20
	*/

	root := &TreeNode{Val: 10}
	root.Left = &TreeNode{Val: 5}
	root.Right = &TreeNode{Val: 15}
	root.Right.Left = &TreeNode{Val: 11}
	root.Right.Right = &TreeNode{Val: 20}
	//root.Left.Left = &TreeNode{Val: 25}

	if isBST(root) {
		fmt.Println("The binary tree is a BST")
	} else {
		fmt.Println("The binary tree is NOT a BST")
	}
}
