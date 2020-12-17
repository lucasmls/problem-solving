package isomorphicstrings

/**
 * Given two strings s and t, determine if they are isomorphic.
 * Two strings are isomorphic if the characters in s can be replaced to get t.
 * All occurrences of a character must be replaced with another character while preserving the order of characters. No two characters may map to the same character but a character may map to itself.
 *
 * Input: s = "egg", t = "add"
 * Output: true
 *
 * Input: s = "foo", t = "bar"
 * Output: false
 *
 * @link https://leetcode.com/problems/isomorphic-strings/
 */

func isIsomorphic(s string, t string) bool {
	sMap := map[string]string{}
	tMap := map[string]string{}

	for i := 0; i < len(s); i++ {
		sLetter := string(s[i])
		tLetter := string(t[i])

		sMappingLetter, ok := sMap[sLetter]
		if !ok {
			sMap[sLetter] = tLetter
		} else {
			if sMappingLetter != tLetter {
				return false
			}
		}

		tMappingLetter, ok := tMap[tLetter]
		if !ok {
			tMap[tLetter] = sLetter
		} else {
			if tMappingLetter != sLetter {
				return false
			}
		}
	}

	return true
}
