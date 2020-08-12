package capitalize

import (
	"fmt"
	"strings"
)

/**
 * Write a function that accepts a string.  The function should
 * capitalize the first letter of each word in the string then
 * return the capitalized string.

 * capitalize('a short sentence') --> 'A Short Sentence'
 * capitalize('a lazy fox') --> 'A Lazy Fox'
 * capitalize('look, it is working!') --> 'Look, It Is Working!'
 */

// Capitalize ...
func Capitalize(phrase string) string {
	words := strings.Fields(phrase)

	var capitalizedWords []string

	for i := 0; i < len(words); i++ {
		word := words[i]

		capitalizedFirstWord := strings.ToUpper(word[:1])
		rest := word[1:]

		capitalizedWords = append(capitalizedWords, fmt.Sprintf("%s%s", capitalizedFirstWord, rest))

	}

	return strings.Join(capitalizedWords, " ")
}
