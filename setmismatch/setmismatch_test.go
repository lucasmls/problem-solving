package setmismatch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindErrorsNums(t *testing.T) {
	tt := []struct {
		input  []int
		output []int
	}{
		{
			input:  []int{1, 2, 2, 4},
			output: []int{2, 3},
		},
		{
			input:  []int{1, 5, 3, 2, 2, 7, 6, 4, 8, 9},
			output: []int{2, 10},
		},
	}

	for _, tc := range tt {
		assert.Equal(t, tc.output, findErrorNums(tc.input))
	}
}
