package sumuptok

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGradingStudents(t *testing.T) {
	testCases := []struct {
		numbersInput []int
		kInput       int
		output       bool
	}{
		{
			numbersInput: []int{10, 15, 3, 7},
			kInput:       17,
			output:       true,
		},
		{
			numbersInput: []int{10, 15, 3, 7},
			kInput:       20,
			output:       false,
		},
		{
			numbersInput: []int{0, 0, 0, 0},
			kInput:       0,
			output:       true,
		},
		{
			numbersInput: []int{0, 0, 0, 0},
			kInput:       1,
			output:       false,
		},
	}

	for _, tc := range testCases {
		result := SumUpToK(tc.numbersInput, tc.kInput)
		assert.Equal(t, tc.output, result)
	}
}
