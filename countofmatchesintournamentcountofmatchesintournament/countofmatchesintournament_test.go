package countofmatchesintournament

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumberOfMatches(t *testing.T) {
	tt := []struct {
		input  int
		output int
	}{
		// #1
		{
			input:  7,
			output: 6,
		},
		// #2
		{
			input:  14,
			output: 13,
		},
		// #3
		{
			input:  21,
			output: 20,
		},
	}

	for _, tc := range tt {
		assert.Equal(t, tc.output, numberOfMatches(tc.input))
	}
}
