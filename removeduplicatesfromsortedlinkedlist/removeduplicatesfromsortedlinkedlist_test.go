package removeduplicatesfromsortedlinkedlist

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
						Val: 12,
						Next: &ListNode{
							Val:  12,
							Next: nil,
						},
					},
				},
			},
			output: &ListNode{
				Val: 10,
				Next: &ListNode{
					Val: 11,
					Next: &ListNode{
						Val:  12,
						Next: nil,
					},
				},
			},
		},
		{
			input: &ListNode{
				Val: 10,
				Next: &ListNode{
					Val: 11,
					Next: &ListNode{
						Val: 12,
						Next: &ListNode{
							Val: 12,
							Next: &ListNode{
								Val: 13,
								Next: &ListNode{
									Val: 13,
									Next: &ListNode{
										Val: 14,
										Next: &ListNode{
											Val: 14,
											Next: &ListNode{
												Val: 15,
												Next: &ListNode{
													Val:  15,
													Next: nil,
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			output: &ListNode{
				Val: 10,
				Next: &ListNode{
					Val: 11,
					Next: &ListNode{
						Val: 12,
						Next: &ListNode{
							Val: 13,
							Next: &ListNode{
								Val: 14,
								Next: &ListNode{
									Val:  15,
									Next: nil,
								},
							},
						},
					},
				},
			},
		},
	}

	for _, tc := range tt {
		result := RemoveDuplicates(tc.input)
		assert.Equal(t, result, tc.output)
	}
}
