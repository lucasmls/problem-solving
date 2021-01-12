package removealladjacentduplicatesinstring

import "github.com/lucasmls/problem-solving/datastructures/stack"

/**
 * Given a string S of lowercase letters, a duplicate removal consists of choosing two adjacent and equal letters, and removing them.
 * We repeatedly make duplicate removals on S until we no longer can.
 * Return the final string after all such duplicate removals have been made.  It is guaranteed the answer is unique.
 *
 * Input: "abbaca"
 * Output: "ca"
 * Explanation:  For example, in "abbaca" we could remove "bb" since the letters are adjacent and equal, and this is the only possible move.
 * The result of this move is that the string is "aaca", of which only "aa" is possible, so the final string is "ca". *
 *
 * @link https://leetcode.com/problems/remove-all-adjacent-duplicates-in-string/
 */

func removeDuplicates(S string) string {
	stack := stack.Stack{}

	for i := 0; i < len(S); i++ {
		letter := string(S[i])

		topStackLetter, err := stack.Peek()
		if err != nil {
			stack.Push(letter)
			continue
		}

		if letter == topStackLetter {
			stack.Pop()
			continue
		}

		stack.Push(letter)
	}

	return stack.ToString()
}
