package isomorphicstrings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsIsomorphic(t *testing.T) {
	tt := []struct {
		firstInput  string
		secondInput string
		output      bool
	}{
		{
			firstInput:  "egg",
			secondInput: "add",
			output:      true,
		},
		{
			firstInput:  "foo",
			secondInput: "bar",
			output:      false,
		},
		{
			firstInput:  "aa",
			secondInput: "ab",
			output:      false,
		},
		{
			firstInput:  "ab",
			secondInput: "aa",
			output:      false,
		},
		{
			firstInput:  "paper",
			secondInput: "title",
			output:      true,
		},
	}

	for _, tc := range tt {
		assert.Equal(t, tc.output, isIsomorphic(tc.firstInput, tc.secondInput))
	}
}
