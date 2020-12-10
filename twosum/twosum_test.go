package twosum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {
	testCases := []struct {
		firstInput  []int
		secondInput int
		output      []int
	}{
		// #1
		{
			firstInput:  []int{10, 15, 3, 7},
			secondInput: 17,
			output:      []int{0, 3},
		},
		// #2
		{
			firstInput:  []int{3, 2, 4},
			secondInput: 6,
			output:      []int{1, 2},
		},
		// #3
		{
			firstInput:  []int{3, 3},
			secondInput: 6,
			output:      []int{0, 1},
		},
	}

	for _, tc := range testCases {
		assert.ElementsMatch(t, tc.output, twoSum(tc.firstInput, tc.secondInput))
	}
}
