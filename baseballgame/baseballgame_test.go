package baseballgame

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalPoints(t *testing.T) {
	tt := []struct {
		input  []string
		output int
	}{
		// #1
		{
			input:  []string{"5", "2", "C", "D", "+"},
			output: 30,
		},
		// #2
		{
			input:  []string{"5", "-2", "4", "C", "D", "9", "+", "+"},
			output: 27,
		},
		// #3
		{
			input:  []string{"1"},
			output: 1,
		},
	}

	for _, tc := range tt {
		assert.Equal(t, tc.output, calPoints(tc.input))
	}
}
