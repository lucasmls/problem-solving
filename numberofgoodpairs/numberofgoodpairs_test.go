package numberofgoodpairs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumOfGoodPairs(t *testing.T) {
	tt := []struct {
		input  []int
		output int
	}{
		// #1
		{
			input:  []int{1, 2, 3, 1, 1, 3},
			output: 4,
		},
		// #2
		{
			input:  []int{1, 1, 1, 1},
			output: 6,
		},
		// #3
		{
			input:  []int{1, 2, 3},
			output: 0,
		},
	}

	for _, tc := range tt {
		assert.Equal(t, tc.output, NumOfGoodPairs(tc.input))
		assert.Equal(t, tc.output, NumOfGoodPairsHM(tc.input))
	}
}
