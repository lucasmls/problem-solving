package longestwordindictionary

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestWord(t *testing.T) {
	tt := []struct {
		input  []string
		output string
	}{
		// #1
		{
			input:  []string{"w", "wo", "wor", "worl", "world"},
			output: "world",
		},
		// #2
		{
			input:  []string{"a", "banana", "app", "appl", "ap", "apply", "apple"},
			output: "apple",
		},
	}

	for _, tc := range tt {
		assert.Equal(t, tc.output, longestWord(tc.input))
	}
}
