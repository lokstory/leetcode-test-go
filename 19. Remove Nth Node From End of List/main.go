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

func wonderfulSolution(head *ListNode, n int) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head
	var length int

	first := head
	for first != nil {
		length++
		first = first.Next
	}

	length -= n
	first = dummy

	for length > 0 {
		length--
		first = first.Next
	}

	first.Next = first.Next.Next
	return dummy.Next
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

	//node := &ListNode{Val: 1}
	//nodes = append(nodes, node)

	//node2 := &ListNode{Val: 2}
	//nodes = append(nodes, node2)
	//
	//node.Next = node2

	//fmt.Println(letterCombinations("23"))
	fmt.Println(removeNthFromEnd(nodes[0], 2))
}
