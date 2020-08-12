package staircase

import (
	"fmt"
	"strings"
)

/* Consider a staircase of size :
 *			 #
 *			##
 *		 ###
 *		####
 *
 *	Observe that its base and height are both equal to , and the image is drawn using # symbols and spaces. The last line is not preceded by any spaces.
 *	Write a program that prints a staircase of size .
 *
 * @link https://www.hackerrank.com/challenges/staircase/problem
 */

// Staircase ...
func Staircase(n int32) {
	staircaseSize := int(n)
	for i := staircaseSize - 1; i >= 0; i-- {
		hashesToPrint := staircaseSize - i
		spacesToPrint := staircaseSize - hashesToPrint

		var spaces strings.Builder
		for i := 0; i < spacesToPrint; i++ {
			fmt.Fprintf(&spaces, " ")
		}

		var hashes strings.Builder
		for i := 0; i < hashesToPrint; i++ {
			fmt.Fprintf(&hashes, "#")
		}

		fmt.Printf("%s%s\n", spaces.String(), hashes.String())
	}
}
