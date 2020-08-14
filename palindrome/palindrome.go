package palindrome

import (
	"fmt"
	"math"
)

/*
 * Given a string, return true if the string is a palindrome
 * or false if it is not.  Palindromes are strings that
 * form the same word if it is reversed. *Do* include spaces
 * and punctuation in determining if the string is a palindrome.
 *  palindrome("abba") === true
 *  palindrome("abcdefg") === false
 */

// Palindrome ...
func Palindrome(str string) bool {
	middle := math.Floor(float64(len(str) / 2))

	for i := 0; i < int(middle); i++ {
		letter := fmt.Sprintf("%c", str[i])
		againstLetter := fmt.Sprintf("%c", str[len(str)-(i+1)])

		if letter != againstLetter {
			return false
		}
	}

	return true
}
