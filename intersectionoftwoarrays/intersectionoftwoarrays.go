package intersectionoftwoarrays

/**
 * Given two arrays, write a function to compute their intersection.
 *
 * Input: nums1 = [1,2,2,1], nums2 = [2,2]
 * Output: [2]
 *
 * Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
 * Output: [9,4]
 *
 * @link https://leetcode.com/problems/intersection-of-two-arrays/
 */

// Intersection ...
func Intersection(nums1 []int, nums2 []int) []int {
	nums1Map := map[int]bool{}
	nums2Map := map[int]bool{}

	for _, num := range nums1 {
		nums1Map[num] = true
	}

	for _, num := range nums2 {
		nums2Map[num] = true
	}

	it := nums1Map
	cmp := nums2Map

	if len(nums2Map) < len(nums1Map) {
		it = nums2Map
		cmp = nums1Map
	}

	result := []int{}
	for num := range it {
		if _, ok := cmp[num]; !ok {
			continue
		}

		result = append(result, num)
	}

	return result
}
