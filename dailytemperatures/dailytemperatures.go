package dailytemperatures

/**
 * Given a list of daily temperatures T, return a list such that, for each day in the input, tells you how many days you would have to wait until a warmer temperature.
 * If there is no future day for which this is possible, put 0 instead.
 * For example, given the list of temperatures
 * T = [73, 74, 75, 71, 69, 72, 76, 73], your output should be [1, 1, 4, 2, 1, 1, 0, 0].
 *
 * @link https://leetcode.com/problems/daily-temperatures/
 */

func dailyTemperatures(T []int) []int {
	result := []int{}

	for i := 0; i < len(T); i++ {
		temp := T[i]
		result = append(result, 0)

		if i+1 < len(T) {
			for j := i + 1; j < len(T); j++ {
				warmerTemp := T[j]

				if warmerTemp > temp {
					result[i] = j - i
					break
				}
			}
		}
	}

	return result
}
