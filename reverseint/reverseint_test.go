package reverseint

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {
	testCases := []struct {
		input  int64
		output int
	}{
		{
			input:  15,
			output: 51,
		},
		{
			input:  981,
			output: 189,
		},
		{
			input:  500,
			output: 5,
		},
		{
			input:  -15,
			output: -51,
		},
		{
			input:  -90,
			output: -9,
		},
	}
	for _, tc := range testCases {
		result := ReverseInt(tc.input)
		assert.Equal(t, tc.output, result)
	}
}
