package longestwordindictionary

import "sort"

/**
 * Given a list of strings words representing an English Dictionary, find the longest word in words that can
 * be built one character at a time by other words in words. If there is more than one possible answer, return
 * the longest word with the smallest lexicographical order.
 *
 * Input: words = ["w","wo","wor","worl", "world"]
 * Output: Output: "world"
 * Explanation: The word "world" can be built one character at a time by "w", "wo", "wor", and "worl".
 *
 * Input: words = ["a", "banana", "app", "appl", "ap", "apply", "apple"]
 * Output: "apple"
 * Explanation: Both "apply" and "apple" can be built from other words in the dictionary. However, "apple" is lexicographically smaller than "apply".
 *
 * @link https://leetcode.com/problems/longest-word-in-dictionary/
 */

func longestWord(words []string) string {
	wordMap := map[string]bool{}
	for _, word := range words {
		wordMap[word] = true
	}

	sort.Slice(words, func(i, j int) bool {
		return words[i] > words[j]
	})

	biggestWord := ""

	for _, word := range words {
		if len(word) < len(biggestWord) {
			break
		}

		for i := 1; i <= len(word); i++ {
			wordSlice := word[0:i]
			_, ok := wordMap[wordSlice]
			if !ok {
				break
			}

			if len(wordSlice) > len(biggestWord) {
				biggestWord = wordSlice
				continue
			}

			if len(wordSlice) == len(biggestWord) && wordSlice < biggestWord {
				biggestWord = wordSlice
			}
		}
	}

	return biggestWord
}
