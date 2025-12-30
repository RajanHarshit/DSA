/*
The idea is to do an iterative level order traversal of the given tree using queue.
If we find a node whose left child is empty, we make a new key as the left child of the node.
Else if we find a node whose right child is empty, we make the new key as the right child.
We keep traversing the tree until we find a node whose either left or right child is empty
*/
package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func insertNode(root *TreeNode, key int) *TreeNode {
	newNode := &TreeNode{Val: key}

	if root == nil {
		return newNode
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		// Insert as Left child
		if node.Left == nil {
			node.Left = newNode
			break
		} else {
			queue = append(queue, node.Left)
		}

		// Insert as Right child
		if node.Right == nil {
			node.Right = newNode
			break
		} else {
			queue = append(queue, node.Right)
		}
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

	key := 7
	root = insertNode(root, key)
	fmt.Println(root.Right.Left)
	if searchNode(root, key) {
		fmt.Println("Node inserted successfully")
	} else {
		fmt.Println("Node insertion failed")
	}

}
