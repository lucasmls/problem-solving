package backspacestringcompare

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBackspaceCompare(t *testing.T) {
	tt := []struct {
		firstInput  string
		secondInput string
		output      bool
	}{
		// #1
		{
			firstInput:  "ab#c",
			secondInput: "ad#c",
			output:      true,
		},
		// #2
		{
			firstInput:  "ab##",
			secondInput: "c#d#",
			output:      true,
		},
		// #3
		{
			firstInput:  "a##c",
			secondInput: "#a#c",
			output:      true,
		},
		// #4
		{
			firstInput:  "a#c",
			secondInput: "b",
			output:      false,
		},
	}

	for _, tc := range tt {
		assert.Equal(t, tc.output, backspaceCompare(tc.firstInput, tc.secondInput))
	}
}
