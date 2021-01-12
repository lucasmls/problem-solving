package backspacestringcompare

import "github.com/lucasmls/problem-solving/datastructures/stack"

/**
 * Given two strings S and T, return if they are equal when both are typed into empty text editors. # means a backspace character.
 * Note that after backspacing an empty text, the text will continue empty.
 *
 * Input: S = "ab#c", T = "ad#c"
 * Output: true
 * Explanation: Both S and T become "ac".
 *
 * Input: S = "a#c", T = "b"
 * Output: false
 * Explanation: S becomes "c" while T becomes "b".
 *
 * @link https://leetcode.com/problems/backspace-string-compare/
 */

func backspaceCompare(S string, T string) bool {
	sStack := stack.Stack{}
	tStack := stack.Stack{}

	for i := 0; i < len(S); i++ {
		char := string(S[i])
		if char == "#" {
			sStack.Pop()
			continue
		}

		sStack.Push(char)
	}

	for i := 0; i < len(T); i++ {
		char := string(T[i])
		if char == "#" {
			tStack.Pop()
			continue
		}

		tStack.Push(char)
	}

	sText := sStack.ToString()
	tText := tStack.ToString()

	if sText == tText {
		return true
	}

	return false
}
