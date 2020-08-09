package disappearednumbers

import "fmt"

/**
 * Given an array of integers where 1 ≤ a[i] ≤ n (n = size of array), some elements appear twice and others appear once.
 * Find all the elements of [1, n] inclusive that do not appear in this array.
 *
 * @link https://leetcode.com/problems/find-all-numbers-disappeared-in-an-array/
 */

// FindDisappearedNumbers ...
func FindDisappearedNumbers(nums []int) []int {
	numbersMap := make(map[int]bool)

	for i := 0; i < len(nums); i++ {
		numbersMap[nums[i]] = true
	}

	var result []int
	for i := 1; i < len(nums)+1; i++ {
		fmt.Println(i)
		if !numbersMap[i] {
			result = append(result, i)
		}
	}

	return result
}
