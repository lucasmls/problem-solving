package subtracttheproductandsumofdigitsofaninteger

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubtractProductAndSum(t *testing.T) {
	tt := []struct {
		input  int
		output int
	}{
		// #1
		{
			input:  234,
			output: 15,
		},
		// #2
		{
			input:  4421,
			output: 21,
		},
	}

	for _, tc := range tt {
		assert.Equal(t, tc.output, subtractProductAndSum(tc.input))
	}
}
