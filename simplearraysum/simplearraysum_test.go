package simplearraysum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimpleArraySum(t *testing.T) {
	testCases := []struct {
		input  []int
		output int
	}{
		{
			input:  []int{5, 10, 15},
			output: 30,
		},
	}
	for _, tc := range testCases {
		result := SimpleArraySum(tc.input)
		assert.Equal(t, tc.output, result)
	}
}
