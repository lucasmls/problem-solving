package pascalstriangle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerate(t *testing.T) {
	testCases := []struct {
		input  int
		output [][]int
	}{
		{
			input: 5,
			output: [][]int{
				{1},
				{1, 1},
				{1, 2, 1},
				{1, 3, 3, 1},
				{1, 4, 6, 4, 1},
			},
		},
		{
			input: 6,
			output: [][]int{
				{1},
				{1, 1},
				{1, 2, 1},
				{1, 3, 3, 1},
				{1, 4, 6, 4, 1},
				{1, 5, 10, 10, 5, 1},
			},
		},
	}
	for _, tC := range testCases {

		result := Generate(tC.input)
		assert.Equal(t, tC.output, result)
	}
}
