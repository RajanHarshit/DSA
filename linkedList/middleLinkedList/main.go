/*
You are given the head of a linked list, You have to return the value of the middle node of the linked list.

If the number of nodes is odd, return the middle node value.
If the number of nodes is even, there are two middle nodes, so return the second middle node value.
*/
package main

import "fmt"

type Node struct {
	data int
	next *Node
}

func length(head *Node) int {
	count := 0
	temp := head
	for temp != nil {
		count++
		temp = temp.next
	}
	return count
}

func getMiddle(head *Node) int {
	if head == nil {
		return -1
	}

	first, second := head, head
	for first != nil && first.next != nil {
		second = second.next
		first = first.next.next
	}
	return second.data
}
func main() {
	head := &Node{data: 1}
	head.next = &Node{data: 2}
	head.next.next = &Node{data: 3}
	head.next.next.next = &Node{data: 4}
	head.next.next.next.next = &Node{data: 5}
	head.next.next.next.next.next = &Node{data: 6}

	fmt.Println(getMiddle(head))
}
