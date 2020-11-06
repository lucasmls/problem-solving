package removeduplicatesfromsortedlinkedlist

/**
 * Given a sorted linked list, delete all duplicates such that each element appear only once.
 *
 * Example 1:
 * Input: 1->1->2
 * Output: 1->2
 *
 * Example 2:
 * Input: 1->1->2->3->3
 * Output: 1->2->3
 *
 * @link https://leetcode.com/problems/remove-duplicates-from-sorted-list/
 */

// ListNode ...
type ListNode struct {
	Val  int
	Next *ListNode
}

// RemoveDuplicates of the given LinkedList
func RemoveDuplicates(head *ListNode) *ListNode {
	var curr *ListNode = head

	for curr != nil && curr.Next != nil {
		if curr.Val == curr.Next.Val {
			curr.Next = curr.Next.Next
			continue
		}

		curr = curr.Next
	}

	return head
}
