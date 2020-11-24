package uncommonwordsfromtwosentences

import "strings"

// UncommonFromSentences ...
func UncommonFromSentences(A string, B string) []string {
	aWords := strings.Split(A, " ")
	bWords := strings.Split(B, " ")

	wordsFrequencyHashMap := map[string]int{}

	for _, word := range aWords {
		if _, ok := wordsFrequencyHashMap[word]; !ok {
			wordsFrequencyHashMap[word] = 0
		}

		wordsFrequencyHashMap[word]++
	}

	for _, word := range bWords {
		if _, ok := wordsFrequencyHashMap[word]; !ok {
			wordsFrequencyHashMap[word] = 0
		}

		wordsFrequencyHashMap[word]++
	}

	result := []string{}
	for word, qtd := range wordsFrequencyHashMap {
		if qtd < 2 {
			result = append(result, word)
		}
	}

	return result
}
