package numberofgoodpairs

/**
 * Given an array of integers nums.
 * A pair (i,j) is called good if nums[i] == nums[j] and i < j.
 * Return the number of good pairs.
 *
 * Input: nums = [1,2,3,1,1,3]
 * Output: 4
 * Explanation: There are 4 good pairs (0,3), (0,4), (3,4), (2,5) 0-indexed.
 *
 * @link https://leetcode.com/problems/number-of-good-pairs/
 */

// NumOfGoodPairs ...
func NumOfGoodPairs(nums []int) int {
	result := 0

	for i, iNum := range nums {
		for j, jNum := range nums {
			if iNum == jNum && i < j {
				result++
			}
		}
	}

	return result
}

// NumOfGoodPairsHM ...
func NumOfGoodPairsHM(nums []int) int {
	result := 0

	hm := make(map[int]int)

	for _, val := range nums {
		_, ok := hm[val]
		if !ok {
			hm[val] = 1
			continue
		}

		hm[val]++

		result += hm[val] - 1
	}

	return result
}
