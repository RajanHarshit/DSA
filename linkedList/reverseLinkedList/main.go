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

func reverseList(head *Node) *Node {
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

func traverse(head *Node) {
	temp := head
	for temp != nil {
		fmt.Print(temp.data, " ")
		temp = temp.next
	}
	fmt.Println()
}

func main() {
	var head *Node

	head = insertAtEnd(head, 1)
	head = insertAtEnd(head, 2)
	head = insertAtEnd(head, 3)
	head = insertAtEnd(head, 4)

	head = reverseList(head)
	traverse(head)
}
