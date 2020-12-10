package longestharmonioussubsequence

import "math"

/**
 * We define a harmonious array as an array where the difference between its maximum value and its minimum value is exactly 1.
 * Given an integer array nums, return the length of its longest harmonious subsequence among all its possible subsequences.
 * A subsequence of array is a sequence that can be derived from the array by deleting some or no elements without changing the order of the remaining elements.
 *
 * Input: nums = [1,3,2,2,5,2,3,7]
 * Output: 5
 * Explanation: The longest harmonious subsequence is [3,2,2,2,3].
 *
 * @link https://leetcode.com/problems/longest-harmonious-subsequence/
 */

func findLHS(nums []int) int {
	hm := map[int]int{}
	for _, num := range nums {
		hm[num]++
	}

	result := 0

	for numKey, numCount := range hm {
		nextNumCount, ok := hm[numKey+1]
		if !ok {
			continue
		}

		result = int(math.Max(float64(result), float64((numCount + nextNumCount))))
	}

	return result
}
