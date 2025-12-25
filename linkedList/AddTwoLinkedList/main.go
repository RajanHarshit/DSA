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

func addLists(head1, head2 *Node) *Node {
	var head, tail *Node
	carry := 0
	for head1 != nil || head2 != nil || carry > 0 {
		sum := carry
		fmt.Printf("sum == carry %d\n ", sum)

		if head1 != nil {
			sum += head1.data
			head1 = head1.next
		}

		fmt.Printf("sum %d \n ", sum)

		if head2 != nil {
			sum += head2.data
			head2 = head2.next
		}
		fmt.Printf("sum %d \n ", sum)
		newNode := &Node{data: sum % 10}
		carry = sum / 10

		if head == nil {
			fmt.Println("if__head__null")
			head = newNode
			fmt.Println("head__", head)
			tail = newNode
			fmt.Println("tail__", tail)
		} else {
			fmt.Println("else__")
			tail.next = newNode
			fmt.Println("tail.next__", tail.next)
			tail = newNode
			fmt.Println("tail", tail)
		}

		fmt.Println()
	}
	return head
}

func removeLeadingZeros(head *Node) *Node {
	for head != nil && head.data == 0 {
		head = head.next
	}
	if head == nil {
		return &Node{data: 0}
	}
	return head
}

func addTwoLists(head1, head2 *Node) *Node {
	// Reverse both lists
	head1 = reverse(head1)
	traverse(head1)
	head2 = reverse(head2)
	traverse(head2)

	// Add reversed lists
	result := addLists(head1, head2)

	// Reverse result
	result = reverse(result)

	// Remove leading zeros
	return removeLeadingZeros(result)
}

func main() {
	var head1, head2 *Node

	// Number 1: 0 → 1 → 2 → 3  (123)
	head1 = insertAtEnd(head1, 0)
	head1 = insertAtEnd(head1, 1)
	head1 = insertAtEnd(head1, 2)
	head1 = insertAtEnd(head1, 3)

	// Number 2: 0 → 9 → 9  (99)
	head2 = insertAtEnd(head2, 0)
	head2 = insertAtEnd(head2, 9)
	head2 = insertAtEnd(head2, 9)

	result := addTwoLists(head1, head2)

	fmt.Print("Result: ")
	traverse(result) // Output: 2 2 2
}
