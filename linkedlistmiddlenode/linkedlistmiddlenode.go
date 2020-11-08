package linkedlistmiddlenode

import (
	"math"
)

/**
 *
 * Given a non-empty, singly linked list with head node head, return a middle node of linked list.
 * If there are two middle nodes, return the second middle node.
 *
 * Input: [1,2,3,4,5]
 * Output: Node 3 from this list (Serialization: [3,4,5])
 *
 * @link https://leetcode.com/problems/middle-of-the-linked-list/
 */

// ListNode ...
type ListNode struct {
	Val  int
	Next *ListNode
}

// MiddleNode ...
func MiddleNode(head *ListNode) *ListNode {
	var middleNode *ListNode = nil
	linkedListLength := 0

	curr := head
	for curr != nil {
		linkedListLength++
		curr = curr.Next
	}

	middlePos := math.Trunc(float64(linkedListLength) / 2)

	curr = head
	for i := 0; float64(i) <= middlePos; i++ {
		if float64(i) == middlePos {
			middleNode = curr
			break
		}
		curr = curr.Next
	}

	return middleNode
}
