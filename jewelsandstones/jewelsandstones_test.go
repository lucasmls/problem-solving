package jewelsandstones

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumJewelsInStones(t *testing.T) {
	testCases := []struct {
		jewelsInput string
		stonesInput string
		output      int
	}{
		{
			jewelsInput: "aA",
			stonesInput: "aAAbbbb",
			output:      3,
		},
		{
			jewelsInput: "a",
			stonesInput: "aAAzzzz",
			output:      1,
		},
	}

	for _, tC := range testCases {
		result := NumJewelsInStones(tC.jewelsInput, tC.stonesInput)
		assert.Equal(t, tC.output, result)
	}
}
