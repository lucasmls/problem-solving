package birthdaycakecandles

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBirthdayCakeCandles(t *testing.T) {
	testCases := []struct {
		input  []int32
		output int32
	}{
		{
			input:  []int32{3, 2, 1, 3},
			output: 2,
		},
		{
			input:  []int32{10, 11, 9, 8, 3},
			output: 1,
		},
	}
	for _, tc := range testCases {
		result := BirthdayCakeCandles(tc.input)
		assert.Equal(t, tc.output, result)
	}
}
