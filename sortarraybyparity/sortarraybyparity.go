package sortarraybyparity

/**
 * Given an array A of non-negative integers, return an array consisting of
 * all the even elements of A, followed by all the odd elements of A.
 *
 * @link https://leetcode.com/problems/sort-array-by-parity
 */

// SortArrayByParity ...
func SortArrayByParity(A []int) []int {
	nextEvenIndex := 0

	for i := 0; i < len(A); i++ {
		actualItem := A[i]
		itemInEvenPosition := A[nextEvenIndex]

		if actualItem%2 == 0 {
			A[nextEvenIndex] = actualItem
			A[i] = itemInEvenPosition
			nextEvenIndex++
		}
	}

	return A
}
