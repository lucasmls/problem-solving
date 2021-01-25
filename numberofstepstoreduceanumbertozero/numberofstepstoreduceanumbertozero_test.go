package numberofstepstoreduceanumbertozero

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumberOfSteps(t *testing.T) {
	tt := []struct {
		input  int
		output int
	}{
		// #1
		{
			input:  14,
			output: 6,
		},
		// #2
		{
			input:  8,
			output: 4,
		},
		// #3
		{
			input:  123,
			output: 12,
		},
	}

	for _, tc := range tt {
		assert.Equal(t, tc.output, numberOfSteps(tc.input))
	}
}
