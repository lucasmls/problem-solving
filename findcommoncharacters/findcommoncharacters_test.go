package findcommoncharacters

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommonChars(t *testing.T) {
	tt := []struct {
		input  []string
		output []string
	}{
		// #1
		{
			input:  []string{"cool", "lock", "cook"},
			output: []string{"c", "o"},
		},
		// #3
		{
			input:  []string{"ll", "ll", "ll"},
			output: []string{"l", "l"},
		},
		// #4
		{
			input:  []string{"l", "l", "l"},
			output: []string{"l"},
		},
		// #5
		{
			input:  []string{},
			output: []string{},
		},
		// #5
		{
			input:  []string{"laaa", "caaa", "baaa"},
			output: []string{"a", "a", "a"},
		},
		// #7
		{
			input:  []string{"acabcddd", "bcbdbcbd", "baddbadb", "cbdddcac", "aacbcccd", "ccccddda", "cababaab", "addcaccd"},
			output: []string{},
		},
	}

	for _, tc := range tt {
		assert.ElementsMatch(t, tc.output, CommonChars(tc.input))
	}
}
