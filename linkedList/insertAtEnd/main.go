package main

import "fmt"

type Node struct {
	data int
	next *Node
}

func insertAtBegin(head *Node, data int) *Node {
	newNode := &Node{
		data: data,
		next: head,
	}
	return newNode
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

func main() {
	var head *Node
	head = insertAtBegin(head, 10)
	head = insertAtEnd(head, 20)
	head = insertAtEnd(head, 30)
	traverse(head)
}
