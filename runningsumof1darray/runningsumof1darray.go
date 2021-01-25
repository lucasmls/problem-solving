package runningsumof1darray

/**
 * Given an array nums. We define a running sum of an array as runningSum[i] = sum(nums[0]…nums[i]).
 * Return the running sum of nums.
 *
 * Input: nums = [1,2,3,4]
 * Output: [1,3,6,10]
 *
 * Input: nums = [1,1,1,1,1]
 * Output: [1,2,3,4,5]
 *
 * @link https://leetcode.com/problems/running-sum-of-1d-array/
 */

func runningSum(nums []int) []int {
	lastSum := 0
	result := []int{}

	for i := 0; i < len(nums); i++ {
		lastSum += nums[i]
		result = append(result, lastSum)
	}

	return result
}
