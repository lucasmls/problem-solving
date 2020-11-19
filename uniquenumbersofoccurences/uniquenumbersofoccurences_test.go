package uniquenumbersofoccurences

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUniqueOccurrences(t *testing.T) {
	tt := []struct {
		input  []int
		output bool
	}{
		// #1
		{
			input:  []int{1, 2, 3, 1, 1, 3},
			output: true,
		},
		// #2
		{
			input:  []int{1, 1, 1, 1},
			output: true,
		},
		// #3
		{
			input:  []int{1, 2, 3},
			output: false,
		},
		// #4
		{
			input:  []int{1, 2, 2, 1, 1, 3},
			output: true,
		},
		// #5
		{
			input:  []int{1, 2},
			output: false,
		},
		// #6
		{
			input:  []int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0},
			output: true,
		},
	}

	for _, tc := range tt {
		assert.Equal(t, tc.output, UniqueOccurrences(tc.input))
	}
}
