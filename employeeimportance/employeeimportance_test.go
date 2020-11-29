package employeeimportance

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetImportance(t *testing.T) {
	tt := []struct {
		firstInput  []*Employee
		secondInput int
		output      int
	}{
		// #0
		{
			firstInput: []*Employee{
				{
					ID:           1,
					Importance:   5,
					Subordinates: []int{2, 3},
				},
				{
					ID:           2,
					Importance:   3,
					Subordinates: []int{},
				},
				{
					ID:           3,
					Importance:   3,
					Subordinates: []int{},
				},
			},
			secondInput: 1,
			output:      11,
		},
	}

	for _, tc := range tt {
		assert.Equal(t, tc.output, GetImportance(tc.firstInput, tc.secondInput))
	}
}
