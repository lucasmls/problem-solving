package shortestcompletingword

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShortestCompletingWord(t *testing.T) {
	tt := []struct {
		firstInput  string
		secondInput []string
		output      string
	}{
		// #1
		{
			firstInput:  "1s3 PSt",
			secondInput: []string{"step", "steps", "stripe", "stepple"},
			output:      "steps",
		},
		// #2
		{
			firstInput:  "1s3 456",
			secondInput: []string{"looks", "pest", "stew", "show"},
			output:      "pest",
		},
		// #3
		{
			firstInput:  "OgEu755",
			secondInput: []string{"enough", "these", "play", "wide", "wonder", "box", "arrive", "money", "tax", "thus"},
			output:      "enough",
		},
		// #4
		{
			firstInput:  "iMSlpe4",
			secondInput: []string{"claim", "consumer", "student", "camera", "public", "never", "wonder", "simple", "thought", "use"},
			output:      "simple",
		},
	}

	for _, tc := range tt {
		assert.Equal(t, tc.output, ShortestCompletingWord(tc.firstInput, tc.secondInput))
	}
}
