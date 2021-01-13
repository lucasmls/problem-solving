package crawlerlogfolder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveDuplicates(t *testing.T) {
	tt := []struct {
		input  []string
		output int
	}{
		// #1
		{
			input:  []string{"d1/", "d2/", "../", "d21/", "./"},
			output: 2,
		},
		// #2
		{
			input:  []string{"d1/", "d2/", "./", "d3/", "../", "d31/"},
			output: 3,
		},
		// #3
		{
			input:  []string{"d1/", "../", "../", "../"},
			output: 0,
		},
	}

	for _, tc := range tt {
		assert.Equal(t, tc.output, minOperations(tc.input))
	}
}
