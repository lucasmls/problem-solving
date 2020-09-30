package productofallnumbers

/**
 * This problem was asked by Uber.

 * Given an array of integers, return a new array such that each element at index i of the new array is
 * the product of all the numbers in the original array except the one at i.
 * For example, if our input was [1, 2, 3, 4, 5], the expected output would be [120, 60, 40, 30, 24].
 * If our input was [3, 2, 1], the expected output would be [2, 3, 6].
 * Follow-up: what if you can't use division?
 */

func Reverse(list []int32) []int32 {
	result := []int32{}
	for i := len(list) - 1; i >= 0; i-- {
		result = append(result, list[i])
	}

	return result
}

// ProductNumbers ...
func ProductNumbers(list []int32) []int32 {
	prefix := []int32{}
	for i := 0; i < len(list); i++ {
		num := list[i]

		if len(prefix) != 0 {
			prefix = append(prefix, prefix[len(prefix)-1]*num)
			continue
		}

		prefix = append(prefix, num)
	}

	reversedList := Reverse(list)

	suffix := []int32{}
	for i := 0; i < len(list); i++ {
		num := reversedList[i]

		if len(suffix) != 0 {
			suffix = append(suffix, suffix[len(suffix)-1]*num)
			continue
		}

		suffix = append(suffix, num)
	}

	suffix = Reverse(suffix)

	products := []int32{}
	for i := 0; i < len(list); i++ {
		if i == 0 {
			products = append(products, suffix[i+1])
			continue
		}

		if i == len(list)-1 {
			products = append(products, prefix[i-1])
			continue
		}

		product := prefix[i-1] * suffix[i+1]
		products = append(products, product)
	}

	return products
}

// ProductNumbers ...
// func ProductNumbers(list []int32) []int32 {
// 	products := []int32{}

// 	var listSum int32 = 1
// 	for i := 0; i < len(list); i++ {
// 		listSum *= list[i]
// 	}

// 	for i := 0; i < len(list); i++ {
// 		result := listSum / list[i]
// 		products = append(products, result)
// 	}

// 	return products
// }
