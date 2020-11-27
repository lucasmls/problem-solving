package maximumnumberofballoons

import "math"

/**
 * Given a string text, you want to use the characters of text to form as many instances of the word "balloon" as possible.
 * You can use each character in text at most once. Return the maximum number of instances that can be formed.
 *
 * Input: text = "nlaebolko"
 * Output: 1
 *
 * Input: text = "loonbalxballpoon"
 * Output: 2
 *
 * @link https://leetcode.com/problems/maximum-number-of-balloons/
 */

// MaxNumberOfBalloons ...
func MaxNumberOfBalloons(text string) int {
	minimumLettersToFormABalloon := map[string]int{
		"b": 1,
		"a": 1,
		"l": 2,
		"o": 2,
		"n": 1,
	}

	inputBalloonLetters := map[string]int{}
	for _, l := range text {
		letter := string(l)

		if _, ok := minimumLettersToFormABalloon[letter]; !ok {
			continue
		}

		if _, ok := inputBalloonLetters[letter]; !ok {
			inputBalloonLetters[letter] = 0
		}

		inputBalloonLetters[letter]++
	}

	qttEachLetterCanFormBalloon := map[string]int{}

	for letter, minQtt := range minimumLettersToFormABalloon {
		qtt, ok := inputBalloonLetters[letter]
		if !ok {
			return 0
		}

		if qtt < minQtt {
			return 0
		}

		qttEachLetterCanFormBalloon[letter] = int(math.Ceil(float64(qtt / minQtt)))
	}

	min := 0
	for _, qtt := range qttEachLetterCanFormBalloon {
		if min == 0 {
			min = qtt
		}

		if qtt < min {
			min = qtt
		}
	}

	return min
}
