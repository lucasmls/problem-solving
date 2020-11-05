package reverselinkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {
	tt := []struct {
		input  *ListNode
		output *ListNode
	}{
		{
			input: &ListNode{
				Val: 10,
				Next: &ListNode{
					Val: 11,
					Next: &ListNode{
						Val:  12,
						Next: nil,
					},
				},
			},
			output: &ListNode{
				Val: 12,
				Next: &ListNode{
					Val: 11,
					Next: &ListNode{
						Val:  10,
						Next: nil,
					},
				},
			},
		},
	}

	for _, tc := range tt {
		result := Reverse(tc.input)
		assert.Equal(t, result, tc.output)
	}
}
