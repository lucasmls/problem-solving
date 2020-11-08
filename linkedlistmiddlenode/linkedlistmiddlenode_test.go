package linkedlistmiddlenode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMiddleNode(t *testing.T) {
	tt := []struct {
		input  *ListNode
		output int
	}{
		{
			input: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 4,
							Next: &ListNode{
								Val: 5,
								Next: &ListNode{
									Val:  6,
									Next: nil,
								},
							},
						},
					},
				},
			},
			output: 4,
		},
		{
			input: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val:  3,
						Next: nil,
					},
				},
			},
			output: 2,
		},
	}

	for _, tc := range tt {
		result := MiddleNode(tc.input)
		assert.Equal(t, tc.output, result.Val)
	}
}
