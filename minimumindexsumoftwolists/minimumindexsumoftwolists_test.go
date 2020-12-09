package minimumindexsumoftwolists

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindRestaurant(t *testing.T) {
	tt := []struct {
		firstInput  []string
		secondInput []string
		output      []string
	}{
		// #1
		{
			firstInput:  []string{"Shogun", "Tapioca Express", "Burger King", "KFC"},
			secondInput: []string{"Piatti", "The Grill at Torrey Pines", "Hungry Hunter Steakhouse", "Shogun"},
			output:      []string{"Shogun"},
		},
		// #2
		{
			firstInput:  []string{"Shogun", "Tapioca Express", "Burger King", "KFC"},
			secondInput: []string{"KFC", "Shogun", "Burger King"},
			output:      []string{"Shogun"},
		},
		// #3
		{
			firstInput:  []string{"Shogun", "Tapioca Express", "Burger King", "KFC"},
			secondInput: []string{"KFC", "Burger King", "Tapioca Express", "Shogun"},
			output:      []string{"Burger King", "KFC", "Shogun", "Tapioca Express"},
		},
		// #4
		{
			firstInput:  []string{"Shogun", "Tapioca Express", "Burger King", "KFC"},
			secondInput: []string{"KNN", "KFC", "Burger King", "Tapioca Express", "Shogun"},
			output:      []string{"Tapioca Express", "Burger King", "KFC", "Shogun"},
		},
	}

	for _, tc := range tt {
		assert.ElementsMatch(t, tc.output, findRestaurant(tc.firstInput, tc.secondInput))
	}
}
