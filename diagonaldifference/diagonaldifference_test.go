package diagonaldifference

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDiagonalDifference(t *testing.T) {
	tt := []struct {
		input  [][]int32
		output int32
	}{
		{
			input: [][]int32{
				{1, 2, 3},
				{4, 5, 6},
				{9, 8, 9},
			},
			output: 2,
		},
		{
			input: [][]int32{
				{10, 10, 1},
				{3, 5, 1},
				{7, 1, 43},
			},
			output: 45,
		},
		{
			input: [][]int32{
				{0, 0, 1},
				{0, 0, 1},
				{0, 0, 1},
			},
			output: 0,
		},
		{
			input: [][]int32{
				{1, 0, 2},
				{0, 1, 1},
				{2, 0, 1},
			},
			output: 2,
		},
	}

	for _, tc := range tt {
		result := DiagonalDifference(tc.input)
		assert.Equal(t, tc.output, result)
	}
}
