package distributecandies

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDistributeCandies(t *testing.T) {
	tt := []struct {
		input  []int
		output int
	}{
		// #1
		{
			input:  []int{1, 1, 2, 2, 3, 3},
			output: 3,
		},
		// #2
		{
			input:  []int{1, 1, 2, 3},
			output: 2,
		},
		// #3
		{
			input:  []int{6, 6, 6, 6},
			output: 1,
		},
		// #4
		{
			input:  []int{},
			output: 0,
		},
	}

	for _, tc := range tt {
		assert.Equal(t, tc.output, DistributeCandies(tc.input))
	}
}
