package intersectionoftwolinkedlists

/**
 * Write a program to find the node at which the intersection of two singly linked lists begins.
 *
 * Input: listA = [4,1,8,4,5], listB = [5,6,1,8,4,5]
 * Output: Reference of the node with value = 8
 *
 * Notes:
 * If the two linked lists have no intersection at all, return null.
 * The linked lists must retain their original structure after the function returns.
 * You may assume there are no cycles anywhere in the entire linked structure.
 * Each value on each linked list is in the range [1, 10^9].
 * Your code should preferably run in O(n) time and use only O(1) memory.
 */

// ListNode is the definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// GetIntersectionNode ...
func GetIntersectionNode(headA, headB *ListNode) *ListNode {
	ht := make(map[*ListNode]int)

	currA := headA

	for currA != nil {
		ht[currA] = currA.Val
		currA = currA.Next
	}

	currB := headB
	for currB != nil {
		_, ok := ht[currB]
		if !ok {
			currB = currB.Next
			continue
		}

		return currB
	}

	return nil
}

// GetIntersectionNodeBruteForce ...
func GetIntersectionNodeBruteForce(headA, headB *ListNode) *ListNode {
	currA := headA
	for currA != nil {

		currB := headB
		for currB != nil {
			if currB == currA {
				return currB
			}

			currB = currB.Next
		}

		currA = currA.Next
	}

	return nil
}
