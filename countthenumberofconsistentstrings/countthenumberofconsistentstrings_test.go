package countthenumberofconsistentstrings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInterpret(t *testing.T) {
	tt := []struct {
		firstInput  string
		secondInput []string
		output      int
	}{
		// #1
		{
			firstInput:  "ab",
			secondInput: []string{"ad", "bd", "aaab", "baa", "badab"},
			output:      2,
		},
		// #2
		{
			firstInput:  "abc",
			secondInput: []string{"a", "b", "c", "ab", "ac", "bc", "abc"},
			output:      7,
		},
		// #3
		{
			firstInput:  "cad",
			secondInput: []string{"cc", "acd", "b", "ba", "bac", "bad", "ac", "d"},
			output:      4,
		},
	}

	for _, tc := range tt {
		assert.Equal(t, tc.output, countConsistentStrings(tc.firstInput, tc.secondInput))
	}
}
