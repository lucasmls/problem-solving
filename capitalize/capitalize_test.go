package capitalize

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	testCases := []struct {
		input  string
		output string
	}{
		// {
		// 	input:  "a short sentence",
		// 	output: "A Short Sentence",
		// },
		{
			input:  "a lazy fox",
			output: "A Lazy Fox",
		},
		{
			input:  "look, it is working!",
			output: "Look, It Is Working!",
		},
	}
	for _, tc := range testCases {
		result := Capitalize(tc.input)
		assert.Equal(t, tc.output, result)
	}
}
