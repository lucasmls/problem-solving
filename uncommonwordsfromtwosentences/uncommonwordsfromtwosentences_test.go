package uncommonwordsfromtwosentences

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUncommonFromSentences(t *testing.T) {
	tt := []struct {
		firstInput  string
		secondInput string
		output      []string
	}{
		// #1
		{
			firstInput:  "this apple is sweet",
			secondInput: "this apple is sour",
			output:      []string{"sweet", "sour"},
		},
		// #2
		{
			firstInput:  "apple apple",
			secondInput: "banana",
			output:      []string{"banana"},
		},
	}

	for _, tc := range tt {
		assert.Equal(t, tc.output, UncommonFromSentences(tc.firstInput, tc.secondInput))
	}
}
