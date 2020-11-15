package linkedlistdeleteelements

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHasCycle(t *testing.T) {
	tt := []struct {
		firstInput  *ListNode
		secondInput int
		output      *ListNode
	}{
		// 1
		{
			firstInput: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val:  4,
							Next: nil,
						},
					},
				},
			},
			secondInput: 3,
			output: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val:  4,
						Next: nil,
					},
				},
			},
		},
		// 2
		{
			firstInput: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 1,
						Next: &ListNode{
							Val:  1,
							Next: nil,
						},
					},
				},
			},
			secondInput: 1,
			output:      nil,
		},
	}

	for _, tc := range tt {
		assert.EqualValues(t, tc.output, RemoveElements(tc.firstInput, tc.secondInput))
	}
}
