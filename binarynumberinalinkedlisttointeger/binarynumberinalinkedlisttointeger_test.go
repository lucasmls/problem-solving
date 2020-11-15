package binarynumberinalinkedlisttointeger

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetDecimalValue(t *testing.T) {
	tt := []struct {
		input  *ListNode
		output int
	}{
		// 1
		{
			input: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 0,
					Next: &ListNode{
						Val: 1,
						Next: &ListNode{
							Val:  1,
							Next: nil,
						},
					},
				},
			},
			output: 11,
		},
		// 2
		{
			input: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 0,
					Next: &ListNode{
						Val: 1,
						Next: &ListNode{
							Val: 0,
							Next: &ListNode{
								Val: 1,
								Next: &ListNode{
									Val:  0,
									Next: nil,
								},
							},
						},
					},
				},
			},
			output: 42,
		},
		// 3
		{
			input: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val:  1,
					Next: nil,
				},
			},
			output: 3,
		},
	}

	for _, tc := range tt {
		assert.Equal(t, tc.output, GetDecimalValue(tc.input))
	}
}
