package sumuptok

/*
 * This problem was recently asked by Google.
 * Given a list of numbers and a number k, return whether any two numbers from the list add up to k.
 * For example, given [10, 15, 3, 7] and k of 17, return true since 10 + 7 is 17.
 * Bonus: Can you do this in one pass?
 */

// SumUpToK ...
func SumUpToK(numbers []int, k int) bool {
	alreadyIteratedNumbers := map[int]bool{}

	for i := 0; i < len(numbers); i++ {
		n := numbers[i]
		expectedNumber := k - n

		if alreadyIteratedNumbers[expectedNumber] {
			return true
		}

		alreadyIteratedNumbers[n] = true
	}

	return false
}
