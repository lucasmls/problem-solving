package reverselinkedlist

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
