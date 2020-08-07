package arraypartition

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArrayPairSum(t *testing.T) {
	tt := []struct {
		nums []int
		want int
	}{
		{
			nums: []int{4, 2, 3, 1},
			want: 4,
		},
	}

	for _, tc := range tt {
		result := ArrayPairSum(tc.nums)
		assert.Equal(t, tc.want, result)
	}
}