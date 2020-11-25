package occurrencesafterbigram

import "strings"

/**
 * Given words first and second, consider occurrences in some text of the form "first second third",
 * where second comes immediately after first, and third comes immediately after second.
 * For each such occurrence, add "third" to the answer, and return the answer.
 *
 * Input: text = "alice is a good girl she is a good student", first = "a", second = "good"
 * Output: ["girl","student"]
 *
 * Input: text = "we will we will rock you", first = "we", second = "will"
 * Output: ["we","rock"]
 *
 * @link https://leetcode.com/problems/occurrences-after-bigram/
 */

// FindOcurrences ...
func FindOcurrences(text string, first string, second string) []string {
	words := strings.Split(text, " ")

	result := []string{}
	for i := 0; i < len(words)-2; i++ {
		if words[i] == first && words[i+1] == second {
			result = append(result, words[i+2])
		}
	}

	return result
}
