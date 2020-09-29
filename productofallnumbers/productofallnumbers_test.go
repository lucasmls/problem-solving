package productofallnumbers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProductNumbers(t *testing.T) {
	tt := []struct {
		input  []int32
		output []int32
	}{
		{
			input:  []int32{3, 2, 1},
			output: []int32{2, 3, 6},
		},
		{
			input:  []int32{1, 2, 3, 4, 5},
			output: []int32{120, 60, 40, 30, 24},
		},
	}

	for _, tc := range tt {
		result := ProductNumbers(tc.input)
		assert.Equal(t, tc.output, result)
	}
}
