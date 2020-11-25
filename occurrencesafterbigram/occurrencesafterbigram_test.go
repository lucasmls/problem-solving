package occurrencesafterbigram

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindOcurrences(t *testing.T) {
	tt := []struct {
		firstInput  string
		secondInput string
		thirdInput  string
		output      []string
	}{
		// #1
		{
			firstInput:  "alice is a good girl she is a good student",
			secondInput: "a",
			thirdInput:  "good",
			output:      []string{"girl", "student"},
		},
		// #2
		{
			firstInput:  "we will we will rock you",
			secondInput: "we",
			thirdInput:  "will",
			output:      []string{"we", "rock"},
		},
	}

	for _, tc := range tt {
		assert.Equal(t, tc.output, FindOcurrences(tc.firstInput, tc.secondInput, tc.thirdInput))
	}
}
