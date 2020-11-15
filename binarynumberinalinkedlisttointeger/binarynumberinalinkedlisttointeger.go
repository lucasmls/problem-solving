package binarynumberinalinkedlisttointeger

import "math"

/**
 * Given head which is a reference node to a singly-linked list. The value of each node in the linked list is either 0 or 1. The linked list holds the binary representation of a number.
 * Return the decimal value of the number in the linked list.
 *
 * Example:
 * Input: head = [1,0,1]
 * Output: 5
 * Explanation: (101) in base 2 = (5) in base 10
 *
 * @link https://leetcode.com/problems/convert-binary-number-in-a-linked-list-to-integer/
 */

// 1*2ˆ3 = 8
// 0*2ˆ2 = 0
// 1*2ˆ1 = 2
// 1*2ˆ0 = 1

// (11)10

// ListNode is the definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// GetDecimalValue ...
func GetDecimalValue(head *ListNode) int {
	size := 0

	i := head
	for i != nil {
		size++
		i = i.Next
	}

	result := 0
	curr := head
	for curr != nil {
		v := curr.Val
		base := 2
		exp := size - 1

		result += int(float64(v) * math.Pow(float64(base), float64(exp)))

		size--
		curr = curr.Next
	}

	return result
}
