package numberofstepstoreduceanumbertozero

/**
 * Given a non-negative integer num, return the number of steps to reduce it to zero.\
 * If the current number is even, you have to divide it by 2, otherwise, you have to subtract 1 from it.
 *
 * Input: num = 14
 * Output: 6
 *
 * @link https://leetcode.com/problems/number-of-steps-to-reduce-a-number-to-zero
 */

func numberOfSteps(num int) int {
	result := 0

	n := num

	for n > 0 {

		if n%2 == 0 {
			n = n / 2
			result++
			continue
		}

		n = n - 1
		result++
	}

	return result
}
