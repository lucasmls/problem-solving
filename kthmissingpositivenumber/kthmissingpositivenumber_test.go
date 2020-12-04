package kthmissingpositivenumber

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindKthPositive(t *testing.T) {
	tt := []struct {
		firstInput  []int
		secondInput int
		output      int
	}{
		// #1
		{
			firstInput:  []int{2, 3, 4, 7, 11},
			secondInput: 5,
			output:      9,
		},
		// #1
		{
			firstInput:  []int{1, 2, 3, 4},
			secondInput: 2,
			output:      6,
		},
	}

	for _, tc := range tt {
		assert.Equal(t, tc.output, findKthPositive(tc.firstInput, tc.secondInput))
	}
}
