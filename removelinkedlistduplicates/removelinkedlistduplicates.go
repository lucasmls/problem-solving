package removelinkedlistduplicates

/**
 * Given any linked list, delete all duplicates such that each element appear only once.
 *
 * Example 1:
 * Input: 1->1->2
 * Output: 1->2
 *
 * Example 2:
 * Input: 4->2->2->10->10
 * Output: 4->2->10
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
	itemsCountHashMap := make(map[int]int)

	var curr *ListNode = head
	var prev *ListNode = nil

	for curr != nil {
		if _, ok := itemsCountHashMap[curr.Val]; !ok {
			itemsCountHashMap[curr.Val] = 0
		}

		if itemsCountHashMap[curr.Val] > 0 {
			prev.Next = curr.Next
			curr = curr.Next
			continue
		}

		prev = curr
		itemsCountHashMap[curr.Val]++

		curr = curr.Next
	}

	return head
}
