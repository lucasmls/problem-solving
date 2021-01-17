package goalparserinterpretation

/**
 * You own a Goal Parser that can interpret a string command.
 * The command consists of an alphabet of "G", "()" and/or "(al)" in some order.
 * The Goal Parser will interpret "G" as the string "G", "()" as the string "o", and "(al)" as the string "al".
 * The interpreted strings are then concatenated in the original order.
 *
 * Input: command = "G()(al)"
 * Output: "Goal"
 *
 * @link https://leetcode.com/problems/goal-parser-interpretation/
 */

func interpret(command string) string {
	result := ""
	i := 0

	for i < len(command) {
		char := string(command[i])

		if char == "G" {
			result += "G"
			i++
			continue
		}

		if char == "(" && string(command[i+1]) == ")" {
			result += "o"
			i += 2
			continue
		}

		if char == "(" && string(command[i+1]) == "a" {
			result += "al"
			i += 3
			continue
		}

		i++
	}

	return result
}
