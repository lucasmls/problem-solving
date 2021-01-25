package richestcustomerwealth

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaximumWealth(t *testing.T) {
	tt := []struct {
		input  [][]int
		output int
	}{
		// #1
		{
			input: [][]int{
				{1, 2, 3},
				{3, 2, 1},
			},
			output: 6,
		},
		// #2
		{
			input: [][]int{
				{1, 5},
				{7, 3},
				{3, 5},
			},
			output: 10,
		},
	}

	for _, tc := range tt {
		assert.Equal(t, tc.output, maximumWealth(tc.input))
	}
}
