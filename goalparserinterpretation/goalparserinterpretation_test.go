package goalparserinterpretation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInterpret(t *testing.T) {
	tt := []struct {
		input  string
		output string
	}{
		// #1
		{
			input:  "G()(al)",
			output: "Goal",
		},
		// #2
		{
			input:  "G()()()()(al)",
			output: "Gooooal",
		},
		// #3
		{
			input:  "(al)G(al)()()G",
			output: "alGalooG",
		},
	}

	for _, tc := range tt {
		assert.Equal(t, tc.output, interpret(tc.input))
	}
}
