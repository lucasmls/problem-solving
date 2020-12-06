package verifyinganaliendictionary

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsAlienSorted(t *testing.T) {
	tt := []struct {
		firstInput  []string
		secondInput string
		output      bool
	}{
		// #1
		{
			firstInput:  []string{"hello", "leetcode"},
			secondInput: "hlabcdefgijkmnopqrstuvwxyz",
			output:      true,
		},
		// #1
		{
			firstInput:  []string{"word", "world", "row"},
			secondInput: "worldabcefghijkmnpqstuvxyz",
			output:      false,
		},
		// #1
		{
			firstInput:  []string{"apple", "app"},
			secondInput: "abcdefghijklmnopqrstuvwxyz",
			output:      false,
		},
	}

	for _, tc := range tt {
		assert.Equal(t, tc.output, isAlienSorted(tc.firstInput, tc.secondInput))
	}
}
