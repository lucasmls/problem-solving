package validparentheses

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsAlienSorted(t *testing.T) {
	tt := []struct {
		input  string
		output bool
	}{
		// #1
		{
			input:  "()",
			output: true,
		},
		// #2
		{
			input:  "()[]{}",
			output: true,
		},
		// #3
		{
			input:  "(]",
			output: false,
		},
		// #4
		{
			input:  "{[]}",
			output: true,
		},
		// #5
		{
			input:  "())",
			output: false,
		},
		// #6
		{
			input:  "(()",
			output: false,
		},
	}

	for _, tc := range tt {
		assert.Equal(t, tc.output, isValidParentheses(tc.input))
	}
}
