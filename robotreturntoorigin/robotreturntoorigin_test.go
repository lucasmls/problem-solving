package robotreturntoorigin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJudgeCircle(t *testing.T) {
	testCases := []struct {
		input  string
		output bool
	}{
		{
			input:  "UD",
			output: true,
		},
		{
			input:  "LL",
			output: false,
		},
		{
			input:  "RLUURDDDLU",
			output: true,
		},
	}
	for _, tC := range testCases {
		result := JudgeCircle(tC.input)
		assert.Equal(t, tC.output, result)
	}
}
