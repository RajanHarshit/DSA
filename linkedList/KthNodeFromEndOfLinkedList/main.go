package main

import "fmt"

type Node struct {
	data int
	next *Node
}

func kthFromEnd(head *Node, k int) int {
	if head == nil || k <= 0 {
		return -1
	}
	first, second := head, head
	for i := 0; i < k; i++ {
		if first == nil {
			return -1 // k > length of list
		}
		first = first.next
	}

	for first != nil {
		first = first.next
		second = second.next
	}
	return second.data

}
func main() {
	head := &Node{data: 1}
	head.next = &Node{data: 2}
	head.next.next = &Node{data: 3}
	head.next.next.next = &Node{data: 4}
	head.next.next.next.next = &Node{data: 5}

	k := 2 //output -> 4
	//k := 6 // output -> -1
	fmt.Println(kthFromEnd(head, k))
}
