package powerfulintegers

import "math"

/**
 * Given two positive integers x and y, an integer is powerful if it is equal to x^i + y^j for some integers i >= 0 and j >= 0.
 * Return a list of all powerful integers that have value less than or equal to bound.
 * You may return the answer in any order.  In your answer, each value should occur at most once.
 *
 * Input: x = 2, y = 3, bound = 10
 * Output: [2,3,4,5,7,9,10]
 *
 * @link https://leetcode.com/problems/powerful-integers/
 */

func powerfulIntegers(x int, y int, bound int) []int {
	hm := map[int]bool{}

	for i := 0; i < 20; i++ {
		for j := 0; j < 20; j++ {
			a := math.Pow(float64(x), float64(i))
			b := math.Pow(float64(y), float64(j))
			sum := int(a + b)

			if sum <= bound && sum > 0 {
				hm[sum] = true
			}
		}
	}

	result := []int{}
	for num := range hm {
		result = append(result, num)
	}

	return result
}
