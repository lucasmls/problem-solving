package minimumaddtomakeparenthesesvalid

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinAddToMakeValid(t *testing.T) {
	tt := []struct {
		input  string
		output int
	}{
		// #1
		{
			input:  "())",
			output: 1,
		},
		// #2
		{
			input:  "(((",
			output: 3,
		},
		// #3
		{
			input:  "()",
			output: 0,
		},
		// #4
		{
			input:  "()))((",
			output: 4,
		},
	}

	for _, tc := range tt {
		assert.Equal(t, tc.output, minAddToMakeValid(tc.input))
	}
}
