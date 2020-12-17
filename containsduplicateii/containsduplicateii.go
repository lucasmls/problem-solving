package containsduplicateii

import "math"

/**
 * Given an array of integers and an integer k, find out whether there are two distinct indices i and j
 * in the array such that nums[i] = nums[j] and the absolute difference between i and j is at most k.
 *
 * Input: nums = [1,2,3,1], k = 3"
 * Output: true
 *
 * Input: nums = [1,2,3,1,2,3], k = 2
 * Output: false
 *
 * @link https://leetcode.com/problems/contains-duplicate-ii/
 */

func containsNearbyDuplicateON(nums []int, k int) bool {
	hm := map[int]int{}

	for i, num := range nums {
		idx, ok := hm[num]
		if !ok {
			idx = -1
		}

		hm[num] = i

		if idx != -1 && i-idx <= k {
			return true
		}
	}

	return false
}

func containsNearbyDuplicate(nums []int, k int) bool {
	for i := 0; i < len(nums)-1; i++ {
		iV := nums[i]

		for j := i + 1; j < len(nums); j++ {
			jV := nums[j]

			if iV == jV && int(math.Abs(float64(i-j))) <= k {
				return true
			}
		}

	}

	return false
}
