package intersectionoftwolinkedlists

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMiddleNode(t *testing.T) {
	intersectionM := &ListNode{
		Val: 9,
		Next: &ListNode{
			Val:  11,
			Next: nil,
		},
	}

	tt := []struct {
		firstInput  *ListNode
		secondInput *ListNode
		output      *ListNode
	}{
		{
			firstInput: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 5,
						Next: &ListNode{
							Val:  7,
							Next: intersectionM,
						},
					},
				},
			},
			secondInput: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val:  4,
					Next: intersectionM,
				},
			},
			output: intersectionM,
		},
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
			secondInput: &ListNode{
				Val: 10,
				Next: &ListNode{
					Val:  11,
					Next: nil,
				},
			},
			output: nil,
		},
	}

	for _, tc := range tt {
		result := GetIntersectionNode(tc.firstInput, tc.secondInput)
		assert.Equal(t, tc.output, result)
	}
}
