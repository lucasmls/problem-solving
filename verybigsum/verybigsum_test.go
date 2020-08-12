package verybigsum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	testCases := []struct {
		input  []int
		output int
	}{
		{
			input: []int{
				1000000001,
				1000000002,
				1000000003,
				1000000004,
				1000000005,
			},
			output: 5000000015,
		},
	}
	for _, tc := range testCases {
		result := VeryBigSum(tc.input)
		assert.Equal(t, tc.output, result)
	}
}
