package nrepeatedelementinsize2narray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRepeatedNTimes(t *testing.T) {
	tt := []struct {
		input  []int
		output int
	}{
		// #1
		{
			input:  []int{1, 2, 3, 3},
			output: 3,
		},
		// #2
		{
			input:  []int{2, 1, 2, 5, 3, 2},
			output: 2,
		},
		// #3
		{
			input:  []int{5, 1, 5, 2, 5, 3, 5, 4},
			output: 5,
		},
		// #3
		{
			input:  []int{0, 0},
			output: 0,
		},
		// #4
		{
			input:  []int{},
			output: 0,
		},
	}

	for _, tc := range tt {
		assert.Equal(t, tc.output, RepeatedNTimes(tc.input))
	}
}
