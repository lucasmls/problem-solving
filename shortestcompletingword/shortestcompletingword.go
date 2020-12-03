package shortestcompletingword

import "strings"

/**
 * Given a string licensePlate and an array of strings words, find the shortest completing word in words.
 * A completing word is a word that contains all the letters in licensePlate. Ignore numbers and spaces in licensePlate, and treat letters as case insensitive. If a letter appears more than once in licensePlate, then it must appear in the word the same number of times or more.
 * For example, if licensePlate = "aBc 12c", then it contains letters 'a', 'b' (ignoring case), and 'c' twice. Possible completing words are "abccdef", "caaacab", and "cbca"
 * Return the shortest completing word in words. It is guaranteed an answer exists. If there are multiple shortest completing words, return the first one that occurs in words.
 *
 * Input: licensePlate = "1s3 PSt", words = ["step","steps","stripe","stepple"]
 * Output: "steps"
 *
 * @link https://leetcode.com/problems/shortest-completing-word/
 */

// ShortestCompletingWord ...
func ShortestCompletingWord(licensePlate string, words []string) string {
	plateMap := map[string]int{}
	for _, letterUnicode := range licensePlate {
		char := string(letterUnicode)

		if letterUnicode >= 65 && letterUnicode <= 90 {
			char = strings.ToLower(char)

			if _, ok := plateMap[char]; !ok {
				plateMap[char] = 0
			}

			plateMap[char]++
		}

		if letterUnicode >= 97 && letterUnicode <= 122 {
			if _, ok := plateMap[char]; !ok {
				plateMap[char] = 0
			}

			plateMap[char]++
		}
	}

	possibleWords := []string{}
	for _, word := range words {
		wordMap := map[string]int{}
		for _, r := range word {
			letter := string(r)
			_, ok := wordMap[letter]
			if !ok {
				wordMap[letter] = 0
			}

			wordMap[letter]++
		}

		inc := true
		for pLetter, pCount := range plateMap {
			wCount, ok := wordMap[pLetter]
			if !ok {
				inc = false
				break
			}

			if pCount > wCount {
				inc = false
				break
			}
		}

		if inc == true {
			possibleWords = append(possibleWords, word)
		}
	}

	minWordLengthIndex := 0
	minWordLength := 16

	for i, word := range possibleWords {
		if len(word) < minWordLength {
			minWordLengthIndex = i
			minWordLength = len(word)
		}
	}

	return possibleWords[minWordLengthIndex]
}
