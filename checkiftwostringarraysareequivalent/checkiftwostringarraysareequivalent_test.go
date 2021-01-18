package checkiftwostringarraysareequivalent

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArrayStringsAreEqual(t *testing.T) {
	tt := []struct {
		firstInput  []string
		secondInput []string
		output      bool
	}{
		// #1
		{
			firstInput:  []string{"ab", "c"},
			secondInput: []string{"a", "bc"},
			output:      true,
		},
		// #1
		{
			firstInput:  []string{"a", "cb"},
			secondInput: []string{"ab", "c"},
			output:      false,
		},
		// #1
		{
			firstInput:  []string{"abc", "d", "defg"},
			secondInput: []string{"abcddefg"},
			output:      true,
		},
	}

	for _, tc := range tt {
		assert.Equal(t, tc.output, arrayStringsAreEqual(tc.firstInput, tc.secondInput))
	}
}
