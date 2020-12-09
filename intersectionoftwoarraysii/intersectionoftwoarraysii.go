package intersectionoftwoarraysii

import (
	"math"
	"sort"
)

/**
 * Given two arrays, write a function to compute their intersection.
 *
 * Input: nums1 = [1,2,2,1], nums2 = [2,2]
 * Output: [2,2]
 *
 * @link https://leetcode.com/problems/intersection-of-two-arrays-ii/
 */

func intersectionOfTwoArrays(nums1 []int, nums2 []int) []int {
	firstMap := map[int]int{}
	for _, num := range nums1 {
		firstMap[num]++
	}

	secondMap := map[int]int{}
	for _, num := range nums2 {
		secondMap[num]++
	}

	result := []int{}
	for num, numCount := range firstMap {
		secondNumCount, ok := secondMap[num]
		if !ok {
			continue
		}

		for i := 0; i < int(math.Min(float64(numCount), float64(secondNumCount))); i++ {
			result = append(result, num)
		}
	}

	return result
}

func intersectionOfTwoArraysTwoPointers(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)

	p1 := 0
	p2 := 0

	result := []int{}

	for p1 < len(nums1) && p2 < len(nums2) {
		if nums1[p1] == nums2[p2] {
			result = append(result, nums1[p1])
			p1++
			p2++
			continue
		}

		if nums1[p1] < nums2[p2] {
			p1++
			continue
		}

		if nums2[p2] < nums1[p1] {
			p2++
			continue
		}
	}

	return result
}
