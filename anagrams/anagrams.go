package anagrams

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

/*
 * Check to see if two provided strings are anagrams of eachother.
 * One string is an anagram of another if it uses the same characters
 * in the same quantity. Only consider characters, not spaces
 * or punctuation.  Consider capital letters to be the same as lower case

 *   anagrams('rail safety', 'fairy tales') --> True
 *   anagrams('RAIL! SAFETY!', 'fairy tales') --> True
 *   anagrams('Hi there', 'Bye there') --> False
 */

func replaceSpecialCharacters(str string) string {
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}

	return reg.ReplaceAllString(str, "")
}

func getStringCharsCountingMap(str string) map[string]int {
	charsCountMap := make(map[string]int)

	for i := 0; i < len(str); i++ {
		char := fmt.Sprintf("%c", str[i])

		if _, ok := charsCountMap[char]; !ok {
			charsCountMap[char] = 1
			continue
		}

		charsCountMap[char]++
	}

	return charsCountMap
}

// Anagrams ...
func Anagrams(stringA string, stringB string) bool {
	firstString := strings.ToLower(replaceSpecialCharacters(stringA))
	secondString := strings.ToLower(replaceSpecialCharacters(stringB))

	firstStringCharsCountingMap := getStringCharsCountingMap(firstString)
	secondStringCharsCountingMap := getStringCharsCountingMap(secondString)

	if len(firstStringCharsCountingMap) != len(secondStringCharsCountingMap) {
		return false
	}

	for letter := range firstStringCharsCountingMap {
		firstStringLetterCount := firstStringCharsCountingMap[letter]
		secondStringLetterCount := secondStringCharsCountingMap[letter]

		if firstStringLetterCount != secondStringLetterCount {
			return false
		}
	}

	return true
}
