package intersectionoftwoarraysii

import (
	"math"
	"sort"
)

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
