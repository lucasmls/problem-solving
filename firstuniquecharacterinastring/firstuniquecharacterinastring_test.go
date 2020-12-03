package firstuniquecharacterinastring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFirstUniqChar(t *testing.T) {
	tt := []struct {
		input  string
		output int
	}{
		// #1
		{
			input:  "leetcode",
			output: 0,
		},
		// #2
		{
			input:  "loveleetcode",
			output: 2,
		},
		// #3
		{
			input:  "aabbcc",
			output: -1,
		},
	}

	for _, tc := range tt {
		assert.Equal(t, tc.output, FirstUniqChar(tc.input))
	}
}
