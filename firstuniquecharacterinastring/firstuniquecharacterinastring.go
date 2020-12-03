package firstuniquecharacterinastring

/**
 * Given a string, find the first non-repeating character in it and return its index. If it doesn't exist, return -1.
 *
 * Input: leetcode
 * Output: 0
 *
 * Input: loveleetcode
 * Output: 2
 *
 * @link https://leetcode.com/problems/first-unique-character-in-a-string/
 */

// FirstUniqChar ...
func FirstUniqChar(s string) int {
	hm := map[string]int{}
	for _, r := range s {
		letter := string(r)
		if _, ok := hm[letter]; !ok {
			hm[letter] = 0
		}

		hm[letter]++
	}

	for i, r := range s {
		letter := string(r)
		count := hm[letter]
		if count == 1 {
			return i
		}
	}

	return -1
}
