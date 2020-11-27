package maximumnumberofballoons

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxNumberOfBalloons(t *testing.T) {
	tt := []struct {
		input  string
		output int
	}{
		// #1
		{
			input:  "nlaebolko",
			output: 1,
		},
		// #2
		{
			input:  "loonbalxballpoon",
			output: 2,
		},
	}

	for _, tc := range tt {
		assert.Equal(t, tc.output, MaxNumberOfBalloons(tc.input))
	}
}
