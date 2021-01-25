package shufflestring

/**
 * Given a string s and an integer array indices of the same length.
 * The string s will be shuffled such that the character at the ith position moves to indices[i] in the shuffled string.
 * Return the shuffled string.
 *
 * Input: s = "abc", indices = [0,1,2]
 * Output: "abc"
 *
 * Input: s = "aiohn", indices = [3,1,4,2,0]
 * Output: "nihao"
 *
 * @link https://leetcode.com/problems/shuffle-string/
 */

func restoreString(s string, indices []int) string {
	shuffledSlice := make([]string, len(s))

	for i := 0; i < len(s); i++ {
		v := string(s[i])
		pos := indices[i]

		shuffledSlice[pos] = v
	}

	result := ""
	for i := 0; i < len(shuffledSlice); i++ {
		result += string(shuffledSlice[i])
	}

	return result
}
