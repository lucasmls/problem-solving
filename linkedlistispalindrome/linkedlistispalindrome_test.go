package linkedlistispalindrome

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPalindrome(t *testing.T) {
	tt := []struct {
		input  *ListNode
		output bool
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
			output: false,
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
			output: false,
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
			output: true,
		},
		// 4
		{
			input: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val:  1,
							Next: nil,
						},
					},
				},
			},
			output: true,
		},
		// 5
		{
			input: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 2,
							Next: &ListNode{
								Val:  1,
								Next: nil,
							},
						},
					},
				},
			},
			output: true,
		},
		// 6
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
								Val:  5,
								Next: nil,
							},
						},
					},
				},
			},
			output: false,
		},
		// 7
		{
			input: &ListNode{
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
			output: false,
		},
		// 7
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
								Val: 4,
								Next: &ListNode{
									Val: 3,
									Next: &ListNode{
										Val: 2,
										Next: &ListNode{
											Val:  1,
											Next: nil,
										},
									},
								},
							},
						},
					},
				},
			},
			output: true,
		},
	}

	for _, tc := range tt {
		assert.Equal(t, tc.output, IsPalindrome(tc.input))
	}
}
