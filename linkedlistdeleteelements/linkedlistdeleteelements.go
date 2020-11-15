package linkedlistdeleteelements

// ListNode is the definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// RemoveElements ...
func RemoveElements(head *ListNode, val int) *ListNode {
	var prev *ListNode = nil
	var curr *ListNode = head

	for curr != nil {
		if curr.Val == val && prev == nil {
			head = curr.Next
			curr = curr.Next
			continue
		}

		if curr.Val == val {
			prev.Next = curr.Next
			curr = curr.Next
			continue
		}

		prev = curr
		curr = curr.Next
	}

	return head
}
