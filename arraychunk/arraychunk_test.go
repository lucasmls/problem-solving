package arraychunk

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChunk(t *testing.T) {
	testCases := []struct {
		listInput []int
		sizeInput int
		output    [][]int
	}{
		{
			listInput: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			sizeInput: 2,
			output: [][]int{
				{1, 2},
				{3, 4},
				{5, 6},
				{7, 8},
				{9, 10},
			},
		},
		{
			listInput: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			sizeInput: 15,
			output: [][]int{
				{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			},
		},
	}
	for _, tc := range testCases {
		result := Chunk(tc.listInput, tc.sizeInput)
		assert.Equal(t, tc.output, result)
	}
}
