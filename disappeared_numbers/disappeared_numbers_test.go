package disappearednumbers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindDisappearedNumbers(t *testing.T) {
	tt := []struct {
		input  []int
		output []int
	}{
		// #1
		{
			input:  []int{4, 3, 2, 7, 8, 2, 3, 1},
			output: []int{5, 6},
		},
		// #2
		{
			input:  []int{1, 3, 5, 7},
			output: []int{2, 4},
		},
		// #3
		{
			input:  []int{1, 1},
			output: []int{2},
		},
	}

	for _, tc := range tt {
		missingNumbers := FindDisappearedNumbers(tc.input)
		assert.Equal(t, tc.output, missingNumbers)
	}
}
