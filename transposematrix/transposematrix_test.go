package transposematrix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTranspose(t *testing.T) {
	tt := []struct {
		matrix [][]int
		want   [][]int
	}{
		// #1
		{
			matrix: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			want: [][]int{
				{1, 4, 7},
				{2, 5, 8},
				{3, 6, 9},
			},
		},

		// #2
		{
			matrix: [][]int{
				{1, 2, 3},
				{4, 5, 6},
			},
			want: [][]int{
				{1, 4},
				{2, 5},
				{3, 6},
			},
		},
	}

	for _, tc := range tt {
		transposed := Transpose(tc.matrix)
		assert.Equal(t, tc.want, transposed)
	}
}
