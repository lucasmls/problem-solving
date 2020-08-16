package maxcommonchar

import "fmt"

/*
 *  Given a string, return the character that is most
 *  commonly used in the string.
 *  maxChar("abcccccccd") === "c"
 *  maxChar("apple 1231111") === "1"
 */

// MostCommonChar ...
type MostCommonChar struct {
	letter   string
	quantity int
}

// MaxCommonChar ...
func MaxCommonChar(str string) string {
	charsCountMap := make(map[string]int)

	for i := 0; i < len(str); i++ {
		char := fmt.Sprintf("%c", str[i])

		if _, ok := charsCountMap[char]; ok {
			charsCountMap[char]++
			continue
		}

		charsCountMap[char] = 1
	}

	mostCommonChar := MostCommonChar{
		letter:   "",
		quantity: 0,
	}

	for char, count := range charsCountMap {
		if count > mostCommonChar.quantity {
			mostCommonChar.letter = char
			mostCommonChar.quantity = count
		}
	}

	return mostCommonChar.letter
}
