package palindrome

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPalindrome(t *testing.T) {
	testCases := []struct {
		input  string
		output bool
	}{
		{
			input:  "abba",
			output: true,
		},
		{
			input:  "abcdefg",
			output: false,
		},
	}
	for _, tc := range testCases {
		result := Palindrome(tc.input)
		assert.Equal(t, tc.output, result)
	}
}
