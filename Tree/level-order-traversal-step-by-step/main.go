package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelorder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	result := [][]int{}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		levelsize := len(queue)
		fmt.Println("level-size__", levelsize)
		level := []int{}

		for i := 0; i < levelsize; i++ {
			node := queue[0]
			queue = queue[1:]
			fmt.Println("queue__", queue)

			level = append(level, node.Val)
			fmt.Println("level__", level)

			if node.Left != nil {
				queue = append(queue, node.Left)
				fmt.Println("if node.Left___", queue)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
				fmt.Println("if node.Right___", queue)
			}
		}

		result = append(result, level)
		fmt.Println("final-result___", result)
		fmt.Println()
		fmt.Println()

	}
	return result
}
func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}
	root.Right = &TreeNode{Val: 3}
	root.Right.Right = &TreeNode{Val: 6}

	fmt.Println(levelorder(root))
}
