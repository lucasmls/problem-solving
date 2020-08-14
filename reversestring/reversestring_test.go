package reversestring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {
	testCases := []struct {
		input string
		output string
	}{
		{
			input:  "apple",
			output: "elppa",
		},
		{
			input:  "hello",
			output: "olleh",
		},
		{
			input:  "Greetings!",
			output: "!sgniteerG",
		},
	}
	for _, tc := range testCases {
		result := Reverse(tc.input)
		assert.Equal(t, tc.output, result)
	}
}
