/*
	You are given the head of a singly linked list.
	You have to determine whether the given linked list contains a loop or not.
	A loop exists in a linked list if the next pointer of the last node points to any other node in the list (including itself),
	rather than being null.

Note: Internally, pos(1 based index) is used to denote the position of the node that tail's next pointer is connected to.
If pos = 0, it means the last node points to null. Note that pos is not passed as a parameter.
*/
package main

type Node struct {
	data int
	next *Node
}

func detectCycle(head *Node) bool {
	if head == nil || head.next == nil {
		return false
	}

	slow, fast := head, head
	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next

		if slow == fast {
			return true
		}
	}
	return false
}

func main() {

}
