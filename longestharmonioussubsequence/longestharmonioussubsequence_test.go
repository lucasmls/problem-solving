package longestharmonioussubsequence

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindLHS(t *testing.T) {
	tt := []struct {
		input  []int
		output int
	}{
		// #1
		{
			input:  []int{1, 3, 2, 2, 5, 2, 3, 7},
			output: 5,
		},
		// #1
		{
			input:  []int{1, 2, 3, 4},
			output: 2,
		},
		// #1
		{
			input:  []int{1, 1, 1, 1},
			output: 0,
		},
	}

	for _, tc := range tt {
		assert.Equal(t, tc.output, findLHS(tc.input))
	}
}
