package powerfulintegers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPowerfulIntegers(t *testing.T) {
	tt := []struct {
		firstInput  int
		secondInput int
		thirdInput  int
		output      []int
	}{
		{
			firstInput:  2,
			secondInput: 3,
			thirdInput:  10,
			output:      []int{2, 3, 4, 5, 7, 9, 10},
		},
		{
			firstInput:  3,
			secondInput: 5,
			thirdInput:  15,
			output:      []int{2, 4, 6, 8, 10, 14},
		},
	}

	for _, tc := range tt {
		assert.ElementsMatch(t, tc.output, powerfulIntegers(tc.firstInput, tc.secondInput, tc.thirdInput))
	}
}
