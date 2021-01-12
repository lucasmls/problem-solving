package validparentheses

import (
	"github.com/lucasmls/problem-solving/datastructures/stack"
)

/**
 * Given a string s containing just the characters '(', ')', '{', '}', '[' and ']',
 * determine if the input string is valid.
 * An input string is valid if:
 * - Open brackets must be closed by the same type of brackets.
 * - Open brackets must be closed in the correct order.
 *
 * Input: s = "()"
 * Output: true
 *
 * Input: s = "()[]{}"
 * Output: true
 *
 * Input: s = "(]"
 * Output: false
 *
 * Input: s = "([)]"
 * Output: false
 *
 * @link https://leetcode.com/problems/valid-parentheses/
 */

func isOpeningBracket(bracket string) bool {
	if bracket == "(" || bracket == "{" || bracket == "[" {
		return true
	}

	return false
}

func bracketsMatchesOpeningAndClosing(opening string, closing string) bool {
	if opening == "(" && closing != ")" {
		return false
	}

	if opening == "{" && closing != "}" {
		return false
	}

	if opening == "[" && closing != "]" {
		return false
	}

	return true
}

func isValidParentheses(s string) bool {
	stack := stack.Stack{}

	for i := 0; i < len(s); i++ {
		bracket := string(s[i])

		if isOpeningBracket(bracket) {
			stack.Push(bracket)
			continue
		}

		opening, err := stack.Peek()
		if err != nil {
			return false
		}

		if !bracketsMatchesOpeningAndClosing(opening, bracket) {
			return false
		}

		stack.Pop()
	}

	if !stack.IsEmpty() {
		return false
	}

	return true
}
