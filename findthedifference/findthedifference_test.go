package findthedifference

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindTheDifference(t *testing.T) {
	tt := []struct {
		firstInput  string
		secondInput string
		output      byte
	}{
		// #1
		{
			firstInput:  "abcd",
			secondInput: "abcde",
			output:      'e',
		},
		// #2
		{
			firstInput:  "",
			secondInput: "y",
			output:      'y',
		},
		// #3
		{
			firstInput:  "a",
			secondInput: "aa",
			output:      'a',
		},
		// #4
		{
			firstInput:  "ae",
			secondInput: "aea",
			output:      'a',
		},
	}

	for _, tc := range tt {
		assert.Equal(t, tc.output, findTheDifference(tc.firstInput, tc.secondInput))
	}
}
