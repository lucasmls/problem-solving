package comparethetriplets

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompareTriplets(t *testing.T) {
	tt := []struct {
		aInput []int32
		bInput []int32
		output []int32
	}{
		{
			aInput: []int32{3, 2, 1},
			bInput: []int32{1, 2, 3},
			output: []int32{1, 1},
		},
		{
			aInput: []int32{5, 6, 7},
			bInput: []int32{3, 6, 10},
			output: []int32{1, 1},
		},
	}

	for _, tc := range tt {
		result := CompareTriplets(tc.aInput, tc.bInput)
		assert.Equal(t, tc.output, result)
	}
}
