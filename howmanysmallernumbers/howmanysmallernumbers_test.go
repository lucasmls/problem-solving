package howmanysmallernumbers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSmallerNumbersThanCurrent(t *testing.T) {
	tt := []struct {
		input  []int
		output []int
	}{
		// #1
		{
			input:  []int{8, 1, 2, 2, 3},
			output: []int{4, 0, 1, 1, 3},
		},
		// #2
		{
			input:  []int{4, 3, 2, 1},
			output: []int{3, 2, 1, 0},
		},
		// #3
		{
			input:  []int{1, 1},
			output: []int{0, 0},
		},
	}

	for _, tc := range tt {
		result := SmallerNumbersThanCurrent(tc.input)
		assert.Equal(t, tc.output, result)
	}
}
