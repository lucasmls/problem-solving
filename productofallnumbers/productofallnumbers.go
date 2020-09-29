package productofallnumbers

/**
 * This problem was asked by Uber.

 * Given an array of integers, return a new array such that each element at index i of the new array is
 * the product of all the numbers in the original array except the one at i.
 * For example, if our input was [1, 2, 3, 4, 5], the expected output would be [120, 60, 40, 30, 24].
 * If our input was [3, 2, 1], the expected output would be [2, 3, 6].
 * Follow-up: what if you can't use division?
 */

// ProductNumbers ...
func ProductNumbers(list []int32) []int32 {
	products := []int32{}

	var listSum int32 = 1
	for i := 0; i < len(list); i++ {
		listSum *= list[i]
	}

	for i := 0; i < len(list); i++ {
		result := listSum / list[i]
		products = append(products, result)
	}

	return products
}
