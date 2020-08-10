package jewelsandstones

import "fmt"

/**
 * You're given strings J representing the types of stones that are jewels, and S representing the
 * stones you have. Each character in S is a type of stone you have. You want to know how many of
 * the stones you have are also jewels.
 *
 * The letters in J are guaranteed distinct, and all characters in J and S are letters.
 * Letters are case sensitive, so "a" is considered a different type of stone from "A".
 *
 * @link https://leetcode.com/problems/jewels-and-stones
 */

// NumJewelsInStones ...
func NumJewelsInStones(J string, S string) int {
	var jewelsMap = make(map[string]bool)
	for _, c := range J {
		char := fmt.Sprintf("%c \n", c)
		jewelsMap[char] = true
	}

	result := 0
	for _, c := range S {
		char := fmt.Sprintf("%c \n", c)
		if _, ok := jewelsMap[char]; ok {
			result++
		}
	}

	return result
}
