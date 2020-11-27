package distributecandies

import "math"

/**
 * Alice has n candies, where the ith candy is of type candyType[i]. Alice noticed that she started to gain weight, o she visited a doctor.
 * The doctor advised Alice to only eat n / 2 of the candies she has (n is always even). Alice likes her candies very much, and she wants to eat the maximum number of different types of candies while still following the doctor's advice.
 * Given the integer array candyType of length n, return the maximum number of different types of candies she can eat if she only eats n / 2 of them.
 *
 * Input: candyType = [1,1,2,2,3,3]
 * Output: 3
 * Explanation: Alice can only eat 6 / 2 = 3 candies. Since there are only 3 types, she can eat one of each type.
 *
 * @link https://leetcode.com/problems/distribute-candies/
 */

// DistributeCandies ...
func DistributeCandies(candyType []int) int {
	if len(candyType) == 0 {
		return 0
	}

	candiesQuantity := map[int]bool{}

	for _, t := range candyType {
		candiesQuantity[t] = true
	}

	maxNum := len(candyType) / 2

	return int(math.Min(float64(maxNum), float64(len(candiesQuantity))))
}
