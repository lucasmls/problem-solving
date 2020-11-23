package findcommoncharacters

/**
 * Given an array A of strings made only from lowercase letters, return a list of all characters
 * that show up in all strings within the list (including duplicates).  For example, if a character
 * occurs 3 times in all strings but not 4 times, you need to include that character three times in
 * the final answer.
 *
 * Input: ["bella","label","roller"]
 * Output: ["e","l","l"]
 *
 * Input: ["cool","lock","cook"]
 * Output: ["c","o"]
 *
 * @link https://leetcode.com/problems/find-common-characters/
 */

// CommonChars ...
func CommonChars(A []string) []string {
	hm := make(map[string][]int)

	wordsCount := len(A)

	for i, word := range A {
		for _, rune := range word {
			letter := string(rune)

			_, ok := hm[letter]
			if !ok {
				hm[letter] = make([]int, wordsCount)
			}

			hm[letter][i]++
		}
	}

	result := []string{}
	for letter, scores := range hm {
		min := scores[0]
		for _, score := range scores {
			if score <= min {
				min = score
			}
		}

		for i := 0; i < min; i++ {
			result = append(result, letter)
		}
	}

	return result
}
