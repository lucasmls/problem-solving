package verifyinganaliendictionary

import (
	"math"
)

/**
 * In an alien language, surprisingly they also use english lowercase letters, but possibly
 * in a different order. The order of the alphabet is some permutation of lowercase letters.
 * Given a sequence of words written in the alien language, and the order of the alphabet,
 * return true if and only if the given words are sorted lexicographicaly in this alien language.
 *
 * Input: words = ["hello","leetcode"], order = "hlabcdefgijkmnopqrstuvwxyz"
 * Output: true
 * Explanation: As 'h' comes before 'l' in this language, then the sequence is sorted.
 *
 * Input: words = ["word","world","row"], order = "worldabcefghijkmnpqstuvxyz"
 * Output: false
 * Explanation: As 'd' comes after 'l' in this language, then words[0] > words[1], hence the sequence is unsorted.
 *
 * @link https://leetcode.com/problems/verifying-an-alien-dictionary/
 */

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
