package runningsumof1darray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRunningSum(t *testing.T) {
	tt := []struct {
		input  []int
		output []int
	}{
		// #1
		{
			input:  []int{1, 2, 3, 4},
			output: []int{1, 3, 6, 10},
		},
		// #2
		{
			input:  []int{1, 1, 1, 1, 1},
			output: []int{1, 2, 3, 4, 5},
		},
		// #2
		{
			input:  []int{3, 1, 2, 10, 1},
			output: []int{3, 4, 6, 16, 17},
		},
	}

	for _, tc := range tt {
		assert.Equal(t, tc.output, runningSum(tc.input))
	}
}
