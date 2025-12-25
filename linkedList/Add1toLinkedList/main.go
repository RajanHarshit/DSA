package main

import "fmt"

type Node struct {
	data int
	next *Node
}

func insertAtEnd(head *Node, data int) *Node {
	newNode := &Node{data: data}
	if head == nil {
		return newNode
	}

	temp := head
	for temp.next != nil {
		temp = temp.next
	}

	temp.next = newNode
	return head
}

func traverse(head *Node) {
	temp := head
	for temp != nil {
		fmt.Print(temp.data, " ")
		temp = temp.next
	}
	fmt.Println()
}

func reverse(head *Node) *Node {
	var prev *Node
	curr := head

	for curr != nil {
		next := curr.next
		curr.next = prev
		prev = curr
		curr = next
	}
	return prev
}

func addOne(head *Node) *Node {
	if head == nil {
		return &Node{data: 1}
	}
	// 1. reverse list
	head = reverse(head)
	carry := 1
	curr := head
	var prev *Node

	// 2. add carry
	for curr != nil && carry > 0 {
		sum := curr.data + carry
		curr.data = sum % 10
		carry = sum / 10

		prev = curr
		curr = curr.next
	}

	// If carry remains
	if carry > 0 {
		prev.next = &Node{data: carry}
	}

	return reverse(head)
}

func main() {
	var head *Node

	// Number: 1 → 9 → 9  (199)
	head = insertAtEnd(head, 1)
	head = insertAtEnd(head, 9)
	head = insertAtEnd(head, 9)

	fmt.Print("Before: ")
	traverse(head)

	head = addOne(head)

	fmt.Print("After:  ")
	traverse(head) // 2 → 0 → 0
}
