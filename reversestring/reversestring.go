package reversestring

import (
	"fmt"
	"strings"
)

/**
 * Given a string, return a new string with the reversed order of characters
 *  reverse("apple") === "elppa";
 *  reverse("hello") === "olleh";
 *  reverse("Greetings!") === "!sgniteerG";
 */

// Reverse ...
func Reverse(str string) string {
	var reversedStr strings.Builder

	for i := len(str) - 1; i >= 0; i-- {
		char := fmt.Sprintf("%c", str[i])
		fmt.Fprintf(&reversedStr, "%s", char)
	}

	return reversedStr.String()
}
