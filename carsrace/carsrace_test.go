package carsrace

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCarsRace(t *testing.T) {
	tt := []struct {
		input  []int
		output []int
	}{
		// #1
		{
			input:  []int{3, 4, 1, 5, -1, 2},
			output: []int{6, 2, 3, 5, 1, 4},
		},
		// #2
		{
			input:  []int{},
			output: []int{},
		},
	}

	for _, tc := range tt {
		assert.Equal(t, tc.output, CarsRace(tc.input))
	}
}
