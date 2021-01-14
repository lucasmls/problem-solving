package minimumaddtomakeparenthesesvalid

import "github.com/lucasmls/problem-solving/datastructures/stack"

/**
 * Given a string S of '(' and ')' parentheses, we add the minimum number of parentheses ( '(' or ')', and in any positions ) so that the resulting parentheses string is valid.
 * Formally, a parentheses string is valid if and only if:
 *   It is the empty string, or
 *   It can be written as AB (A concatenated with B), where A and B are valid strings, or
 *   It can be written as (A), where A is a valid string.
 * Given a parentheses string, return the minimum number of parentheses we must add to make the resulting string valid.
 *
 * Input: "())"
 * Output: 1
 *
 * Input: "((("
 * Output: 3
 *
 * Input: "()))(("
 * Output: 4
 *
 * @link https://leetcode.com/problems/minimum-add-to-make-parentheses-valid/
 */

func minAddToMakeValid(S string) int {
	result := 0
	stack := stack.Stack{}

	for i := 0; i < len(S); i++ {
		parentheses := string(S[i])

		if parentheses == "(" {
			stack.Push(parentheses)
			continue
		}

		if stack.IsEmpty() {
			result++
			continue
		}

		stack.Pop()
	}

	result += stack.Size()

	return result
}
