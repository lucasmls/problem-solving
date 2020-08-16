package anagrams

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnagrams(t *testing.T) {
	testCases := []struct {
		firstInput  string
		secondInput string
		output      bool
	}{
		{
			firstInput:  "hello",
			secondInput: "llohe",
			output:      true,
		},
		{
			firstInput:  "Whoa! Hi!",
			secondInput: "Hi! Whoa!",
			output:      true,
		},
		{
			firstInput:  "One One",
			secondInput: "Two two two",
			output:      false,
		},
		{
			firstInput:  "One one",
			secondInput: "One one c",
			output:      false,
		},
		{
			firstInput:  "A tree, a life, a bench",
			secondInput: "A tree, a fence, a yard",
			output:      false,
		},
	}
	for _, tc := range testCases {
		result := Anagrams(tc.firstInput, tc.secondInput)
		assert.Equal(t, tc.output, result)
	}
}
