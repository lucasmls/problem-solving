package nrepeatedelementinsize2narray

/**
 * In a array A of size 2N, there are N+1 unique elements, and exactly one of these elements is repeated N times.
 *
 * Input: [2,1,2,5,3,2]
 * Output: 2
 *
 * Input: [5,1,5,2,5,3,5,4]
 * Output: 5
 *
 * @link https://leetcode.com/problems/number-of-good-pairs/
 */

// RepeatedNTimes ...
func RepeatedNTimes(A []int) int {
	hm := make(map[int]int)

	for _, v := range A {
		_, ok := hm[v]
		if !ok {
			hm[v] = 1
			continue
		}

		return v
	}

	return 0
}
