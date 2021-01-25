package shufflestring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRestoreString(t *testing.T) {
	tt := []struct {
		firstInput  string
		secondInput []int
		output      string
	}{
		// #1
		{
			firstInput:  "codeleet",
			secondInput: []int{4, 5, 6, 7, 0, 2, 1, 3},
			output:      "leetcode",
		},
		// #2
		{
			firstInput:  "abc",
			secondInput: []int{0, 1, 2},
			output:      "abc",
		},
		// #3
		{
			firstInput:  "aiohn",
			secondInput: []int{3, 1, 4, 2, 0},
			output:      "nihao",
		},
		// #4
		{
			firstInput:  "aaiougrt",
			secondInput: []int{4, 0, 2, 6, 7, 3, 1, 5},
			output:      "arigatou",
		},
		// #5
		{
			firstInput:  "art",
			secondInput: []int{1, 0, 2},
			output:      "rat",
		},
	}

	for _, tc := range tt {
		assert.Equal(t, tc.output, restoreString(tc.firstInput, tc.secondInput))
	}
}
