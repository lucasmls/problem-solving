package keyboardrow

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindWords(t *testing.T) {
	tt := []struct {
		input  []string
		output []string
	}{
		// #1
		{
			input:  []string{"Hello", "Alaska", "Dad", "Peace"},
			output: []string{"Alaska", "Dad"},
		},
	}

	for _, tc := range tt {
		assert.Equal(t, tc.output, FindWords(tc.input))
	}
}
