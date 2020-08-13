package plusminus

import "fmt"

/**
 * Given an array of integers, calculate the ratios of its elements that are
 * positive, negative, and zero. Print the decimal value of each fraction on a
 * new line with  places after the decimal.

 * Example
 * arr = [1, 1, 0, -1, -1]
 * There are n = 5 elements, two positive, two negative and one zero.
 * Their ratios are 2/5 = 0.400000, 2/5 = 0.400000 and 1/5 - 0.200000. Results are printed as:
 * | 0.400000
 * | 0.400000
 * | 0.200000
 */

// PlusMinus ...
func PlusMinus(arr []int32) {
	var positives int32
	var negatives int32
	var zeros int32

	for i := 0; i < len(arr); i++ {
		el := arr[i]

		if el >= 1 {
			positives++
		}

		if el == 0 {
			zeros++
		}

		if el <= -1 {
			negatives++
		}
	}

	itemsCount := float64(len(arr))

	positiveResult := fmt.Sprintf("%.6f", float64(positives)/itemsCount)
	negativeResult := fmt.Sprintf("%.6f", float64(negatives)/itemsCount)
	zeroResult := fmt.Sprintf("%.6f", float64(zeros)/itemsCount)

	fmt.Println(positiveResult)
	fmt.Println(negativeResult)
	fmt.Println(zeroResult)

}
