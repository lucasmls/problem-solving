package dailytemperatures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDailyTemperatures(t *testing.T) {
	tt := []struct {
		input  []int
		output []int
	}{
		// #1
		{
			input:  []int{73, 74, 75, 71, 69, 72, 76, 73},
			output: []int{1, 1, 4, 2, 1, 1, 0, 0},
		},
		// #2
		{
			input:  []int{70, 70, 70, 75},
			output: []int{3, 2, 1, 0},
		},
	}

	for _, tc := range tt {
		assert.Equal(t, tc.output, dailyTemperatures(tc.input))
	}
}
