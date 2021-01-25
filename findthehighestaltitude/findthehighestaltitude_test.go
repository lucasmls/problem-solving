package findthehighestaltitude

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLargestAltitude(t *testing.T) {
	tt := []struct {
		input  []int
		output int
	}{
		// #1
		{
			input:  []int{-5, 1, 5, 0, -7},
			output: 1,
		},
		// #2
		{
			input:  []int{-4, -3, -2, -1, 4, 3, 2},
			output: 0,
		},
	}

	for _, tc := range tt {
		assert.Equal(t, tc.output, largestAltitude(tc.input))
	}
}
