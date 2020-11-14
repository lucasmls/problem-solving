package linkedlistcycle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHasCycle(t *testing.T) {
	nodeOne := &ListNode{Val: 1}
	nodeTwo := &ListNode{Val: 2}
	nodeThree := &ListNode{Val: 3}

	nodeOne.Next = nodeTwo
	nodeTwo.Next = nodeThree
	nodeThree.Next = nodeOne

	tt := []struct {
		input  *ListNode
		output bool
	}{
		// #1
		{
			input:  nodeOne,
			output: true,
		},
		// #2
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
			output: false,
		},
	}

	for _, tc := range tt {
		assert.Equal(t, tc.output, HasCycle(tc.input))
		assert.Equal(t, tc.output, HasCycleTwoPointers(tc.input))
	}
}
