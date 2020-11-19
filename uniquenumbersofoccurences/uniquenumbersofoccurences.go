package uniquenumbersofoccurences

/**
 * Given an array of integers arr, write a function that returns true if and only if the number of occurrences of each value in the array is unique.
 *
 * Input: arr = [1,2,2,1,1,3]
 * Output: true
 * Explanation: The value 1 has 3 occurrences, 2 has 2 and 3 has 1. No two values have the same number of occurrences.
 *
 * @link https://leetcode.com/problems/unique-number-of-occurrences/
 */

// UniqueOccurrences ...
func UniqueOccurrences(arr []int) bool {
	hm := make(map[int]int)
	for _, val := range arr {
		_, ok := hm[val]
		if !ok {
			hm[val] = 1
			continue
		}

		hm[val]++
	}

	hmCount := make(map[int]int)
	for num, qtd := range hm {
		_, ok := hmCount[qtd]
		if !ok {
			hmCount[qtd] = num
			continue
		}

		return false
	}

	return true
}
