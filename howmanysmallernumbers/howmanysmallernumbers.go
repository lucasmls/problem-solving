package howmanysmallernumbers

import (
	"sort"
)

/**
 *
 * Given the array nums, for each nums[i] find out how many numbers in the array are smaller than it.
 * That is, for each nums[i] you have to count the number of valid j's such that j != i and nums[j] < nums[i].
 *
 * @link https://leetcode.com/problems/how-many-numbers-are-smaller-than-the-current-number/
 */

// SmallerNumbersThanCurrent ...
func SmallerNumbersThanCurrent(nums []int) []int {
	var sortedNums []int
	for _, value := range nums {
		sortedNums = append(sortedNums, value)
	}

	sort.Ints(sortedNums)

	numsMap := make(map[int]int)

	for i, value := range sortedNums {
		if _, ok := numsMap[value]; !ok {
			numsMap[value] = i
		}
	}

	var result []int
	for _, value := range nums {
		v := numsMap[value]
		result = append(result, v)
	}

	return result
}
