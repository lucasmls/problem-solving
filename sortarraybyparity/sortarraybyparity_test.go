package sortarraybyparity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	testCases := []struct {
		input  []int
		output []int
	}{
		{
			input:  []int{3, 1, 2, 4},
			output: []int{2, 4, 3, 1},
		},
		{
			input:  []int{1, 6, 3, 8, 11},
			output: []int{6, 8, 3, 1, 11},
		},
	}
	for _, tC := range testCases {
		result := SortArrayByParity(tC.input)
		assert.Equal(t, tC.output, result)
	}
}
