package grading

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGradingStudents(t *testing.T) {
	testCases := []struct {
		input  []int32
		output []int32
	}{
		{
			input:  []int32{73, 67, 38, 33},
			output: []int32{75, 67, 40, 33},
		},
		{
			input:  []int32{38, 37, 36, 33, 32, 31},
			output: []int32{40, 37, 36, 33, 32, 31},
		},
	}

	for _, tc := range testCases {
		result := GradingStudents(tc.input)
		assert.Equal(t, tc.output, result)
	}
}
