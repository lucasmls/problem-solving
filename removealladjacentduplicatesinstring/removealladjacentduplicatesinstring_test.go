package removealladjacentduplicatesinstring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveDuplicates(t *testing.T) {
	tt := []struct {
		input  string
		output string
	}{
		// #1
		{
			input:  "leeetcode",
			output: "letcode",
		},
		// #2
		{
			input:  "abbacc",
			output: "",
		},
	}

	for _, tc := range tt {
		assert.Equal(t, tc.output, removeDuplicates(tc.input))
	}
}
