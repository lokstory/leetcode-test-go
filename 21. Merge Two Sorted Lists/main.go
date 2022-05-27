package main

import (
	"fmt"
)

// You are given the heads of two sorted linked lists list1 and list2.
// Merge the two lists in a one sorted list.
// The list should be made by splicing together the nodes of the first two lists.
// Return the head of the merged linked list.

// Input: list1 = [1,2,4], list2 = [1,3,4]
// Output: [1,1,2,3,4,4]
//
// Input: list1 = [], list2 = []
// Output: []
//
// Input: list1 = [], list2 = [0]
// Output: [0]
//
// Constraints:
// The number of nodes in both lists is in the range [0, 50].
// -100 <= Node.val <= 100
// Both list1 and list2 are sorted in non-decreasing order.

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var ret *ListNode
	var current *ListNode
	var item *ListNode

	for list1 != nil || list2 != nil {
		isFirst := list1 != nil && (list2 == nil || list1.Val <= list2.Val)

		if isFirst {
			item = list1
			list1 = list1.Next
		} else {
			item = list2
			list2 = list2.Next
		}

		if ret == nil {
			ret = item
		} else {
			current.Next = item
		}

		current = item
	}

	return ret
}

func addNode(values []int) *ListNode {
	var nodes []*ListNode
	count := len(values)

	for i := 0; i < count; i++ {
		item := &ListNode{
			Val: values[i],
		}

		nodes = append(nodes, item)

		if i > 0 {
			nodes[i-1].Next = item
		}
	}

	if count > 0 {
		return nodes[0]
	}

	return nil
}

func main() {
	fmt.Println(mergeTwoLists(addNode([]int{1, 2, 4}), addNode([]int{1, 3, 4})))
	fmt.Println(mergeTwoLists(addNode([]int{}), addNode([]int{})))
	fmt.Println(mergeTwoLists(addNode([]int{-10, -10, -9, -4, 1, 6, 6}), addNode([]int{-7})))
}
