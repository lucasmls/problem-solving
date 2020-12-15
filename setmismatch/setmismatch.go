package setmismatch

/**
 * The set S originally contains numbers from 1 to n. But unfortunately, due to the data error, one of the numbers
 * in the set got duplicated to another number in the set, which results in repetition of one number and loss of another number.
 * Given an array nums representing the data status of this set after the error. Your task is to firstly find the number occurs
 * twice and then find the number that is missing. Return them in the form of an array.
 *
 * Input: nums = [1,2,2,4]
 * Output: [2,3]
 *
 * @link https://leetcode.com/problems/set-mismatch/
 */

func findErrorNums(nums []int) []int {
	var duplicatedNumber int
	var missingNumber int

	hm := map[int]int{}
	for i := 1; i <= len(nums); i++ {
		hm[i] = 0
	}

	for _, num := range nums {
		hm[num]++
	}

	for i := 1; i <= len(nums); i++ {
		count := hm[i]

		if count > 1 {
			duplicatedNumber = i
		}

		if count == 0 {
			missingNumber = i
		}
	}

	return []int{duplicatedNumber, missingNumber}
}
