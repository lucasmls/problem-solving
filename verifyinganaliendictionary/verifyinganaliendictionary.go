package verifyinganaliendictionary

import (
	"math"
)

func isAlienSorted(words []string, order string) bool {
	orderMap := map[string]int{}
	for i, r := range order {
		letter := string(r)
		orderMap[letter] = i
	}

	for i := 0; i < len(words)-1; i++ {
		cWord := words[i]
		nWord := words[i+1]

		cWordLen := float64(len(cWord))
		nWordLen := float64(len(nWord))

		compareLength := true
		for k := 0; k < int(math.Min(cWordLen, nWordLen)); k++ {
			cLetter := string(cWord[k])
			nLetter := string(nWord[k])

			if orderMap[cLetter] != orderMap[nLetter] {
				if orderMap[cLetter] > orderMap[nLetter] {
					return false
				}

				compareLength = false
				break
			}
		}

		if compareLength == true && cWordLen > nWordLen {
			return false
		}
	}

	return true
}
