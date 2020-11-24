package keyboardrow

/**
 * Given a List of words, return the words that can be typed using letters of alphabet on only one row's of American keyboard like the image below.
 *
 * Input: ["Hello", "Alaska", "Dad", "Peace"]
 * Output: ["Alaska", "Dad"]
 *
 * @link https://leetcode.com/problems/keyboard-row/
 */

import "strings"

// FindWords ...
func FindWords(words []string) []string {
	rows := []map[string]bool{
		{"q": true, "w": true, "e": true, "r": true, "t": true, "y": true, "u": true, "i": true, "o": true, "p": true},
		{"a": true, "s": true, "d": true, "f": true, "g": true, "h": true, "j": true, "k": true, "l": true},
		{"z": true, "x": true, "c": true, "v": true, "b": true, "n": true, "m": true},
	}

	result := []string{}
	for _, word := range words {
		count := 0

		for _, row := range rows {
			for _, l := range word {
				letter := strings.ToLower(string(l))

				if _, ok := row[letter]; !ok {
					count++
					break
				}
			}
		}

		if count < 3 {
			result = append(result, word)
		}
	}

	return result
}
