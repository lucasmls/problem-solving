package checkiftwostringarraysareequivalent

/**
 * Given two string arrays word1 and word2, return true if the two arrays represent the same string, and false otherwise.
 * A string is represented by an array if the array elements concatenated in order forms the string.
 *
 * Input: word1 = ["ab", "c"], word2 = ["a", "bc"]
 * Output: true
 *
 * @link https://leetcode.com/problems/check-if-two-string-arrays-are-equivalent/
 */

func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	firstString := ""
	secondString := ""

	for i := 0; i < len(word1); i++ {
		word := word1[i]

		for j := 0; j < len(word); j++ {
			char := string(word[j])
			firstString += char
		}
	}

	for i := 0; i < len(word2); i++ {
		word := word2[i]

		for j := 0; j < len(word); j++ {
			char := string(word[j])
			secondString += char
		}
	}

	if firstString == secondString {
		return true
	}

	return false
}
