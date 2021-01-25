package defanginganipaddress

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsDefangIPaddr(t *testing.T) {
	tt := []struct {
		input  string
		output string
	}{
		// #1
		{
			input:  "1.1.1.1",
			output: "1[.]1[.]1[.]1",
		},
		// #2
		{
			input:  "255.100.50.0",
			output: "255[.]100[.]50[.]0",
		},
	}

	for _, tc := range tt {
		assert.Equal(t, tc.output, defangIPaddr(tc.input))
	}
}
