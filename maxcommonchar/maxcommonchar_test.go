package maxcommonchar

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxCommonChar(t *testing.T) {
	testCases := []struct {
		input  string
		output string
	}{
		{
			input:  "abcccccccd",
			output: "c",
		},
		{
			input:  "apple 1231111",
			output: "1",
		},
	}
	for _, tc := range testCases {
		result := MaxCommonChar(tc.input)
		assert.Equal(t, tc.output, result)
	}
}
