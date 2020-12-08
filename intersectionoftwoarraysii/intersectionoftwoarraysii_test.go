package intersectionoftwoarraysii

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntersectionOfTwoArrays(t *testing.T) {
	tt := []struct {
		firstInput  []int
		secondInput []int
		output      []int
	}{
		// #1
		{
			firstInput:  []int{1, 2, 2, 1},
			secondInput: []int{2, 2},
			output:      []int{2, 2},
		},
		// #2
		{
			firstInput:  []int{4, 9, 5},
			secondInput: []int{9, 4, 9, 8, 4},
			output:      []int{4, 9},
		},
	}

	for _, tc := range tt {
		assert.Equal(t, tc.output, intersectionOfTwoArrays(tc.firstInput, tc.secondInput))
		assert.Equal(t, tc.output, intersectionOfTwoArraysTwoPointers(tc.firstInput, tc.secondInput))
	}
}
