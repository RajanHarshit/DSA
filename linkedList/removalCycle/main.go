/*
	Given the head of a singly linked list, the task is to remove a cycle if present.
	A cycle exists when a node's next pointer points back to a previous node, forming a loop.
	Internally, a variable pos denotes the index of the node where the cycle starts, but it is not passed as a parameter.
	The terminal will print true if a cycle is removed otherwise, it will print false.

Examples:

Input: head = 1 -> 3 -> 4, pos = 2
Output: true
Explanation: The linked list looks like

A loop is present in the list, and it is removed.
Input: head = 1 -> 8 -> 3 -> 4, pos = 0
Output: true
Explanation:

The Linked list does not contains any loop.
Input: head = 1 -> 2 -> 3 -> 4, pos = 1
Output: true
Explanation: The linked list looks like

A loop is present in the list, and it is removed.
*/
package main

import "fmt"

type Node struct {
	data int
	next *Node
}

func removeCycle(head *Node) bool {
	if head == nil || head.next == nil {
		return false
	}

	slow, fast := head, head

	// step 1. Detect cycle
	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
		if slow == fast {
			break
		}
	}

	// no-cycle
	if fast == nil || fast.next == nil {
		return false
	}

	// step 2. start cycle
	slow = head
	for slow != fast {
		slow = slow.next
		fast = fast.next
	}

	// remove cycle
	ptr := slow
	for ptr.next != slow {
		ptr = ptr.next
	}
	ptr.next = nil
	return true
}

func main() {
	head := &Node{data: 1}
	head.next = &Node{data: 2}
	head.next.next = &Node{data: 3}
	head.next.next.next = &Node{data: 4}
	head.next.next.next.next = &Node{data: 5}

	// Create cycle at position 1 (0-based)
	head.next.next.next.next.next = head.next

	fmt.Println(removeCycle(head))
}
