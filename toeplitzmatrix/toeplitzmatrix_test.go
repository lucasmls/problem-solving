package toeplitzmatrix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArrayPairSum(t *testing.T) {
	tt := []struct {
		input  [][]int
		output bool
	}{
		{
			input: [][]int{
				{1, 2, 3, 4},
				{5, 1, 2, 3},
				{9, 5, 1, 2},
			},
			output: true,
		},
		{
			input: [][]int{
				{1, 2, 3, 4},
				{5, 2, 2, 3},
				{9, 5, 1, 2},
			},
			output: false,
		},
	}

	for _, tc := range tt {
		result := isToeplitzMatrix(tc.input)
		assert.Equal(t, tc.output, result)
	}
}
