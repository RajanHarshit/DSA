/*
The idea is to traverse the tree in level order manner. To perform the Deletion in a Binary Tree follow below:
Starting at the root, find the deepest and rightmost node in the binary tree and the node which we want to delete.
Replace the deepest rightmost node’s data with the node to be deleted.
Then delete the deepest rightmost node.
*/
package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}

	// Single Node Tree
	if root.Left == nil && root.Right == nil {
		if root.Val == key {
			return nil
		}
		return root
	}
	var targetNode *TreeNode
	var lastNode *TreeNode
	var parentOfLast *TreeNode

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node.Val == key {
			// write logic to delete node
			targetNode = node

		}
		if node.Left != nil {
			parentOfLast = node
			lastNode = node.Left
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			parentOfLast = node
			lastNode = node.Right
			queue = append(queue, node.Right)
		}
	}

	// if key not found
	if targetNode == nil {
		return root
	}
	// replace value
	targetNode.Val = lastNode.Val

	// delete deepest node
	if parentOfLast.Right == lastNode {
		parentOfLast.Right = nil
	} else {
		parentOfLast.Left = nil
	}
	return root
}

func searchNode(root *TreeNode, key int) bool {
	if root == nil {
		return false
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node.Val == key {
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

	key := 6
	root = deleteNode(root, key)
	fmt.Println(root.Right.Right)
	if searchNode(root, key) {
		fmt.Println("Node deleted successfully")
	} else {
		fmt.Println("Node insertion failed")
	}
}
