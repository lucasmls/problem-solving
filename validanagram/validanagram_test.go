package validanagram

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsAnagram(t *testing.T) {
	tt := []struct {
		firstInput  string
		secondInput string
		output      bool
	}{
		// #1
		{
			firstInput:  "anagram",
			secondInput: "nagaram",
			output:      true,
		},
		// #2
		{
			firstInput:  "rat",
			secondInput: "car",
			output:      false,
		},
		// #3
		{
			firstInput:  "lucas",
			secondInput: "sacul",
			output:      true,
		},
		// #4
		{
			firstInput:  "lucas",
			secondInput: "lala",
			output:      false,
		},
		// #5
		{
			firstInput:  "lucas",
			secondInput: "llucas",
			output:      false,
		},
	}

	for _, tc := range tt {
		assert.Equal(t, tc.output, IsAnagram(tc.firstInput, tc.secondInput))
	}

}
