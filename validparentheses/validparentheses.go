package validparentheses

import "errors"

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

// Stack ...
type Stack struct {
	items []string
}

// Push a new item in to the Stack
func (s *Stack) Push(value string) {
	s.items = append(s.items, value)
}

// Pop out the last inserted item from the Stack
func (s *Stack) Pop() (string, error) {
	if len(s.items) > 0 {
		lastItem := s.items[len(s.items)-1]
		s.items = s.items[:len(s.items)-1]

		return lastItem, nil
	}

	return "", errors.New("The stack is empty")
}

// Peek the last inserted item from the Stack
func (s *Stack) Peek() (string, error) {
	if len(s.items) > 0 {
		return s.items[len(s.items)-1], nil
	}

	return "", errors.New("The stack is empty")
}

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
	stack := Stack{}

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

	if len(stack.items) > 0 {
		return false
	}

	return true
}
