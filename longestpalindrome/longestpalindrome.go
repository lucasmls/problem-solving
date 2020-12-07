package longestpalindrome

/**
 * Given a string s which consists of lowercase or uppercase letters, return the length of the longest palindrome that can be built with those letters.
 * Letters are case sensitive, for example, "Aa" is not considered a palindrome here.
 *
 * Input: abccccdd
 * Output: 7
 * Explanation: One longest palindrome that can be built is "dccaccd", whose length is 7.
 *
 * Input: a
 * Output: 1
 *
 * @link https://leetcode.com/problems/longest-palindrome/
 */

func divide(v int) (int, int) {
	div := v / 2
	rest := v % 2

	return div, rest
}

func longestPalindrome(s string) int {
	lettersMap := map[rune]int{}
	for _, r := range s {
		lettersMap[r]++
	}

	result := 0
	uniqLettersCount := 0

	for _, count := range lettersMap {
		pairs, rest := divide(count)

		result += pairs * 2

		if rest != 0 {
			uniqLettersCount++
		}
	}

	if uniqLettersCount >= 1 {
		result++
	}

	return result
}
