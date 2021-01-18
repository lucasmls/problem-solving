package countthenumberofconsistentstrings

/**
 * You are given a string allowed consisting of distinct characters and an array of strings words.
 * A string is consistent if all characters in the string appear in the string allowed.
 * The command consists of an alphabet of "G", "()" and/or "(al)" in some order.
 * Return the number of consistent strings in the array words.
 *
 * Input: allowed = "ab", words = ["ad","bd","aaab","baa","badab"]
 * Output: 2
 * Explanation: Strings "aaab" and "baa" are consistent since they only contain characters 'a' and 'b'.
 *
 * @link https://leetcode.com/problems/count-the-number-of-consistent-strings/
 */

func countConsistentStrings(allowed string, words []string) int {
	result := 0

	allowedSet := map[string]bool{}
	for i := 0; i < len(allowed); i++ {
		allowedSet[string(allowed[i])] = true
	}

	for i := 0; i < len(words); i++ {
		word := string(words[i])
		result++

		for j := 0; j < len(word); j++ {
			letter := string(word[j])

			if _, ok := allowedSet[letter]; !ok {
				result--
				break
			}
		}
	}

	return result
}
