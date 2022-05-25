package main

import (
	"fmt"
)

// Given a linked list, remove the n-th node from the end of list and return its head.
//
// Example:
//
// Given linked list: 1->2->3->4->5, and n = 2.
//
// After removing the second node from the end, the linked list becomes 1->2->3->5.

// Note:
//
// Given n will always be valid.

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	nodeMap := map[int]*ListNode{}
	var startID int
	var counter int

	current := head
	for current != nil {
		nodeMap[startID+counter] = current
		counter++

		current = current.Next
	}

	removeID := startID + counter - n
	prev, ok := nodeMap[removeID-1]
	next, _ := nodeMap[removeID+1]

	if ok {
		prev.Next = next
		return head
	}

	return next
}

func main() {
	var nodes []*ListNode
	for i := 1; i <= 2; i++ {
		node := &ListNode{Val: i}
		nodes = append(nodes, node)

		if i > 1 {
			nodes[i-1-1].Next = node
		}
	}

	fmt.Println(removeNthFromEnd(nodes[0], 2))
}
