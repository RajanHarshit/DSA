package main

import "fmt"

type Node struct {
	data int
	next *Node
}

func isPalindrome(head *Node) bool {
	if head == nil || head.next == nil {
		return true
	}

	// find middle
	slow, fast := head, head
	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
	}

	first := head
	// reverse second half
	secondHalf := reverseList(slow)
	for secondHalf != nil {
		if first.data != secondHalf.data {
			return false
		}
		first = first.next
		secondHalf = secondHalf.next
	}

	return true

}

func reverseList(head *Node) *Node {
	var prev *Node
	curr := head
	//fmt.Println("curr__", curr)

	for curr != nil {
		next := curr.next
		fmt.Println("next__", next)
		curr.next = prev
		fmt.Println("curr.next", curr.next)
		prev = curr
		fmt.Println("prev", prev)
		curr = next
		fmt.Println("curr", curr)

	}

	return prev
}

func insertAtEnd(head *Node, val int) *Node {
	newNode := &Node{data: val}
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

func main() {
	var head *Node

	// Create linked list: 1 → 2 → 3 → 2 → 1
	head = insertAtEnd(head, 1)
	head = insertAtEnd(head, 2)
	head = insertAtEnd(head, 3)
	head = insertAtEnd(head, 2)
	head = insertAtEnd(head, 1)

	if isPalindrome(head) {
		fmt.Println("Linked List is Palindrome")
	} else {
		fmt.Println("Linked List is NOT Palindrome")
	}
}
