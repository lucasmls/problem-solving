package arraypartition

import (
	"sort"
)

/**
 * Given an array of 2n integers, your task is to group these integers into n pairs of integer,
 * say (a1, b1), (a2, b2), ..., (an, bn) which makes sum of min(ai, bi) for all i from 1 to n as large as possible.
 * Example:
 * Input: [1,4,3,2]
 * Output: 4
 *
 * @link https://leetcode.com/problems/array-partition-i/
 */

// ArrayPairSum ...
func ArrayPairSum(nums []int) int {
	sort.Ints(nums)

	result := 0
	for i := 0; i < len(nums); i += 2 {
		result += nums[i]
	}

	return result
}
