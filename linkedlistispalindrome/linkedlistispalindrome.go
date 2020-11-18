package linkedlistispalindrome

/**
 *
 * Given a singly linked list, determine if it is a palindrome.
 *
 * Input: 1->2->2->1
 * Output: true
 *
 * @link https://leetcode.com/problems/palindrome-linked-list/
 */

// ListNode is the definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// IsPalindrome ...
func IsPalindrome(head *ListNode) bool {
	var slowPtr *ListNode = head
	var fastPtr *ListNode = head

	for fastPtr != nil && fastPtr.Next != nil {
		fastPtr = fastPtr.Next.Next
		slowPtr = slowPtr.Next
	}

	if fastPtr != nil {
		slowPtr = slowPtr.Next
	}

	secHead := slowPtr

	var ch *ListNode = secHead
	var chPrev *ListNode = nil
	for ch != nil {
		tempNext := ch.Next
		ch.Next = chPrev
		chPrev = ch
		ch = tempNext
	}

	a := chPrev
	b := head
	for a != nil {
		if a.Val != b.Val {
			return false
		}

		a = a.Next
		b = b.Next
	}

	return true
}
