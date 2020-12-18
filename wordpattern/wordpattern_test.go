package wordpattern

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWordPattern(t *testing.T) {
	tt := []struct {
		firstInput  string
		secondInput string
		output      bool
	}{
		// #1
		{
			firstInput:  "abba",
			secondInput: "dog cat cat dog",
			output:      true,
		},
		// #2
		{
			firstInput:  "abba",
			secondInput: "dog cat cat fish",
			output:      false,
		},
		// #3
		{
			firstInput:  "aaaa",
			secondInput: "dog cat cat dog",
			output:      false,
		},
		// #3
		{
			firstInput:  "abbba",
			secondInput: "dog cat cat fish",
			output:      false,
		},
		{
			firstInput:  "abbba",
			secondInput: "dog dog dog dog",
			output:      false,
		},
	}

	for _, tc := range tt {
		assert.Equal(t, tc.output, wordPattern(tc.firstInput, tc.secondInput))
	}
}
