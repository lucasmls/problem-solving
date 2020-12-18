package wordpattern

import (
	"strings"
)

/**
 * Given a pattern and a string s, find if s follows the same pattern.
 * Here follow means a full match, such that there is a bijection between a letter in pattern and a non-empty word in s.
 *
 * Input: pattern = "abba", s = "dog cat cat dog"
 * Output: true
 *
 * Input: pattern = "abba", s = "dog cat cat fish"
 * Output: false
 *
 * @link https://leetcode.com/problems/word-pattern/
 */

func wordPattern(pattern string, s string) bool {
	words := strings.Split(s, " ")

	wordsMap := map[string]string{}
	letterMap := map[string]string{}

	if len(words) != len(pattern) {
		return false
	}

	for i, wordByte := range words {
		word := string(wordByte)
		letter := string(pattern[i])

		wordMapLetter, ok := wordsMap[word]
		if !ok {
			wordsMap[word] = letter
		} else {
			if wordMapLetter != letter {
				return false
			}
		}

		letterMapWord, ok := letterMap[letter]
		if !ok {
			letterMap[letter] = word
		} else {
			if letterMapWord != word {
				return false
			}
		}
	}

	return true
}
