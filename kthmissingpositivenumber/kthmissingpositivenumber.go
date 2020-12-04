package kthmissingpositivenumber

/**
 * Given an array arr of positive integers sorted in a strictly increasing order, and an integer k.
 * Find the kth positive integer that is missing from this array.
 *
 * Input: arr = [2,3,4,7,11], k = 5
 * Output: 9
 * Explanation: The missing positive integers are [1,5,6,8,9,10,12,13,...]. The 5th missing positive integer is 9.
 *
 * Input: arr = [1,2,3,4], k = 2
 * Output: 6
 * Explanation: The missing positive integers are [5,6,7,...]. The 2nd missing positive integer is 6.
 *
 * @link https://leetcode.com/problems/kth-missing-positive-number/
 */

func findKthPositive(arr []int, k int) int {
	hm := map[int]bool{}
	for _, n := range arr {
		hm[n] = true
	}

	missingNumbers := []int{}
	i := 1
	for len(missingNumbers) < k {
		if _, ok := hm[i]; !ok {
			missingNumbers = append(missingNumbers, i)
		}

		i++
	}

	return missingNumbers[k-1]
}

func findKthPositiveBinarySearch(arr []int, k int) int {
	left := 0
	right := len(arr)
	mid := 0

	for left < right {
		mid = (left + right) / 2

		if (arr[mid] - 1 - mid) < k {
			left = mid + 1
			continue
		}

		right = mid
	}

	return left + k
}
