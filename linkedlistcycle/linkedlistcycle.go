package linkedlistcycle

/**
 * Given head, the head of a linked list, determine if the linked list has a cycle in it.
 * Return true if there is a cycle in the linked list. Otherwise, return false.
 *
 * @link https://leetcode.com/problems/linked-list-cycle/
 */

// ListNode is the definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// HasCycle ...
func HasCycle(head *ListNode) bool {
	ht := make(map[*ListNode]interface{})

	curr := head
	for curr != nil {
		_, contains := ht[curr]
		if !contains {
			ht[curr] = true
			curr = curr.Next
			continue
		}

		return true
	}

	return false
}

// HasCycleTwoPointers ...
func HasCycleTwoPointers(head *ListNode) bool {
	slow := head
	fast := head.Next

	for slow != fast {
		if fast == nil || fast.Next == nil {
			return false
		}

		slow = slow.Next
		fast = fast.Next.Next
	}

	return true
}
