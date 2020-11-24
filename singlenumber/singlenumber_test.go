package singlenumber

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSingleNumber(t *testing.T) {
	tt := []struct {
		input  []int
		output int
	}{
		// #1
		{
			input:  []int{2, 2, 1},
			output: 1,
		},
		// #2
		{
			input:  []int{4, 1, 2, 1, 2},
			output: 4,
		},
		// #3
		{
			input:  []int{1},
			output: 1,
		},
	}

	for _, tc := range tt {
		assert.Equal(t, tc.output, SingleNumber(tc.input))
	}
}
