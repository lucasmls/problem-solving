package validanagram

/**
 * Given two strings s and t , write a function to determine if t is an anagram of s.
 *
 * Input: s = "anagram", t = "nagaram"
 * Output: true
 *
 * Input: s = "rat", t = "car"
 * Output: false
 *
 * @link https://leetcode.com/problems/valid-anagram/
 */

// IsAnagram ...
func IsAnagram(s string, t string) bool {
	sMap := map[string]int{}
	for _, char := range s {
		key := string(char)
		if _, ok := sMap[key]; !ok {
			sMap[key] = 0
		}

		sMap[key]++
	}

	tMap := map[string]int{}
	for _, char := range t {
		key := string(char)
		if _, ok := tMap[key]; !ok {
			tMap[key] = 0
		}

		tMap[key]++
	}

	itMap := sMap
	compMap := tMap
	if len(sMap) != len(tMap) {
		return false
	}

	for char, count := range itMap {
		compCount, ok := compMap[char]
		if !ok {
			return false
		}

		if count != compCount {
			return false
		}
	}

	return true
}
