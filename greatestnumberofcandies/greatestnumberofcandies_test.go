package greatestnumberofcandies

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	testCases := []struct {
		candies      []int
		extraCandies int
		output       []bool
	}{
		{
			candies:      []int{2, 3, 5, 1, 3},
			extraCandies: 3,
			output:       []bool{true, true, true, false, true},
		},
		{
			candies:      []int{1, 10},
			extraCandies: 6,
			output:       []bool{false, true},
		},
	}
	for _, tC := range testCases {
		result := KidsWithCandies(tC.candies, tC.extraCandies)
		assert.Equal(t, tC.output, result)
	}
}
