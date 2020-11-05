package removelinkedlistduplicates

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
