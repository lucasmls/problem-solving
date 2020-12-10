package twosum

/**
 * Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
 * You may assume that each input would have exactly one solution, and you may not use the same element twice.
 *
 * Input: nums = [2,7,11,15], target = 9
 * Output: [0,1]
 * Output: Because nums[0] + nums[1] == 9, we return [0, 1].
 *
 * Input: nums = [3,2,4], target = 6
 * Output: [1,2]
 *
 * @link https://leetcode.com/problems/two-sum/
 */

func twoSum(nums []int, target int) []int {
	hm := map[int]int{}

	for i, num := range nums {
		desiredValue := target - num

		desiredIndex, ok := hm[desiredValue]
		if !ok {
			hm[num] = i
			continue
		}

		return []int{i, desiredIndex}
	}

	return []int{}
}
