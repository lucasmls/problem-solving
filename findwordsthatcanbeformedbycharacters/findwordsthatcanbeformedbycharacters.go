package findwordsthatcanbeformedbycharacters

/**
 * You are given an array of strings words and a string chars.
 * A string is good if it can be formed by characters from chars (each character can only be used once).
 * Return the sum of lengths of all good strings in words.
 *
 * Input: words = ["cat","bt","hat","tree"], chars = "atach"
 * Output: 6
 * Explanation: The strings that can be formed are "cat" and "hat" so the answer is 3 + 3 = 6.
 *
 * Input: words = ["hello","world","leetcode"], chars = "welldonehoneyr"
 * Output: 10
 * Explanation: The strings that can be formed are "hello" and "world" so the answer is 5 + 5 = 10.
 *
 * @link https://leetcode.com/problems/find-words-that-can-be-formed-by-characters/
 */

// CountCharacters ...
func CountCharacters(words []string, chars string) int {
	charsHM := map[string]int{}

	for _, letterRune := range chars {
		letter := string(letterRune)
		if _, ok := charsHM[letter]; !ok {
			charsHM[letter] = 0
		}

		charsHM[letter]++
	}

	result := 0
	for _, word := range words {
		lettersMap := map[string]int{}
		for k, v := range charsHM {
			lettersMap[k] = v
		}

		include := true

		for _, letterRune := range word {
			letter := string(letterRune)

			_, ok := lettersMap[letter]
			if !ok {
				include = false
				break
			}

			if lettersMap[letter] <= 0 {
				include = false
				break
			}

			lettersMap[letter]--
		}

		if include == true {
			result += len(word)
		}
	}

	return result
}
