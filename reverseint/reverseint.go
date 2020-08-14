package reverseint

import (
	"fmt"
	"strconv"
)

/*
 * Given an integer, return an integer that is the reverse
 * ordering of numbers.

 *   reverseInt(15) === 51
 *   reverseInt(981) === 189
 *   reverseInt(500) === 5
 *   reverseInt(-15) === -51
 *   reverseInt(-90) === -9
 */

// ReverseInt ...
func ReverseInt(n int64) int {
	str := strconv.FormatInt(n, 10)

	stringifiedResult := ""
	for i := len(str) - 1; i >= 0; i-- {
		char := fmt.Sprintf("%c", str[i])

		if char == "0" {
			continue
		}

		if char == "-" {
			stringifiedResult = fmt.Sprintf("%s%s", char, stringifiedResult)
			continue
		}

		stringifiedResult = fmt.Sprintf("%s%s", stringifiedResult, char)
	}

	r, _ := strconv.Atoi(stringifiedResult)
	return r
}
