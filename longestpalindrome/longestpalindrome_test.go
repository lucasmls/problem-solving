package longestpalindrome

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestPalindrome(t *testing.T) {
	tt := []struct {
		input  string
		output int
	}{
		{
			input:  "abccccdd",
			output: 7,
		},
		{
			input:  "a",
			output: 1,
		},
		{
			input:  "bb",
			output: 2,
		},
	}

	for _, tc := range tt {
		assert.Equal(t, tc.output, longestPalindrome(tc.input))
	}
}
