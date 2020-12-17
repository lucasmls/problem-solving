package containsduplicateii

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContainsNearbyDuplicate(t *testing.T) {
	tt := []struct {
		firstInput  []int
		secondInput int
		output      bool
	}{
		// #1
		{
			firstInput:  []int{1, 2, 3, 1},
			secondInput: 3,
			output:      true,
		},
		// #2
		{
			firstInput:  []int{1, 0, 1, 1},
			secondInput: 1,
			output:      true,
		},
		// #3
		{
			firstInput:  []int{1, 2, 3, 1, 2, 3},
			secondInput: 2,
			output:      false,
		},
	}

	for _, tc := range tt {
		assert.Equal(t, tc.output, containsNearbyDuplicateON(tc.firstInput, tc.secondInput))
		assert.Equal(t, tc.output, containsNearbyDuplicate(tc.firstInput, tc.secondInput))
	}
}
