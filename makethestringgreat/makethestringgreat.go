package makethestringgreat

import (
	"strings"

	"github.com/lucasmls/problem-solving/datastructures/stack"
)

/**
 * Given a string s of lower and upper case English letters.
 * A good string is a string which doesn't have two adjacent characters s[i] and s[i + 1] where:
 * 	-	0 <= i <= s.length - 2
 * 	-	s[i] is a lower-case letter and s[i + 1] is the same letter but in upper-case or vice-versa.
 *
 * To make the string good, you can choose two adjacent characters that make the string bad and remove them.
 * You can keep doing this until the string becomes good.
 * Return the string after making it good. The answer is guaranteed to be unique under the given constraints.
 *
 * Input: s = "leEeetcode"
 * Output: "leetcode"
 * Explanation: In the first step, either you choose i = 1 or i = 2, both will result "leEeetcode" to be reduced to "leetcode".
 *
 * Input: s = "abBAcC"
 * Output: ""
 * Explanation: We have many possible scenarios, and all lead to the same answer. For example:
 *  - "abBAcC" --> "aAcC" --> "cC" --> ""
 *  - "abBAcC" --> "abBA" --> "aA" --> ""
 *
 * @link https://leetcode.com/problems/make-the-string-great/
 */

func makeGood(s string) string {
	stack := stack.Stack{}

	for i := 0; i < len(s); i++ {
		letter := string(s[i])

		topStackLetter, err := stack.Peek()
		if err != nil {
			stack.Push(letter)
			continue
		}

		if letter != topStackLetter && strings.ToLower(letter) == strings.ToLower(topStackLetter) {
			stack.Pop()
			continue
		}

		stack.Push(letter)
	}

	return stack.ToString()
}
