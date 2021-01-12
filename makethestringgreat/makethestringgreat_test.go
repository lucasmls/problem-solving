package makethestringgreat

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakeGood(t *testing.T) {
	tt := []struct {
		input  string
		output string
	}{
		// #1
		{
			input:  "leEeetcode",
			output: "leetcode",
		},
		// #2
		{
			input:  "abBAcC",
			output: "",
		},
	}

	for _, tc := range tt {
		assert.Equal(t, tc.output, makeGood(tc.input))
	}
}
