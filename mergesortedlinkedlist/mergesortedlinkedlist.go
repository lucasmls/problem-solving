package mergesortedlinkedlist

/**
 *
 * Merge two sorted linked lists and return it as a new sorted list. The new list should be made by splicing together the nodes of the first two lists.
 *
 * Input: l1 = [1,2,4], l2 = [1,3,4]
 * Output: [1,1,2,3,4,4]
 *
 * Input: l1 = [], l2 = [0]
 * Output: [0]
 *
 * @link https://leetcode.com/problems/merge-two-sorted-lists/
 */

// ListNode ...
type ListNode struct {
	Val  int
	Next *ListNode
}

// MergeSortedLinkedList ...
func MergeSortedLinkedList(l1 *ListNode, l2 *ListNode) *ListNode {
	var sortedArray []int = []int{}

	var c1 *ListNode = l1
	var c2 *ListNode = l2

	for c1 != nil && c2 != nil {
		if c1.Val <= c2.Val {
			sortedArray = append(sortedArray, c1.Val)
			c1 = c1.Next
			continue
		}

		sortedArray = append(sortedArray, c2.Val)
		c2 = c2.Next
	}

	for c1 != nil {
		sortedArray = append(sortedArray, c1.Val)
		c1 = c1.Next
	}

	for c2 != nil {
		sortedArray = append(sortedArray, c2.Val)
		c2 = c2.Next
	}

	var result *ListNode = nil
	for _, v := range sortedArray {
		if result == nil {
			result = &ListNode{Val: v, Next: nil}
			continue
		}

		cr := result
		for cr.Next != nil {
			cr = cr.Next
		}

		cr.Next = &ListNode{Val: v, Next: nil}
	}

	return result
}
