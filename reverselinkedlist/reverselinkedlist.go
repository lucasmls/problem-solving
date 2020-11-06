package reverselinkedlist

/**
 * Reverse a singly linked list.
 *
 * Example 1:
 * Input: 1->2->3->4->5->NULL
 * Output: 5->4->3->2->1->NULL
 *
 * @link https://leetcode.com/problems/reverse-linked-list/
 */

// ListNode ...
type ListNode struct {
	Val  int
	Next *ListNode
}

// Reverse the given LinkedList
func Reverse(head *ListNode) *ListNode {
	var prev *ListNode = nil
	var curr *ListNode = head

	for curr != nil {
		tempNext := curr.Next
		curr.Next = prev

		prev = curr
		curr = tempNext
	}

	return prev
}
