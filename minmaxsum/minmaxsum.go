package minmaxsum

import (
	"fmt"
	"sort"
)

/**
 * Given five positive integers, find the minimum and maximum values that can be calculated by summing exactly four of the five integers. Then print the respective minimum and maximum values as a single line of two space-separated long integers.

 * Example
 * The minimum sum is  and the maximum sum is . The function prints
 * 16 24

 * @link https://www.hackerrank.com/challenges/mini-max-sum/problem
 */

// MinMaxSum ...
func MinMaxSum(arr []int32) {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	firstFour := arr[:4]
	lastFour := arr[1:]

	var firstSum int64
	for i := 0; i < len(firstFour); i++ {
		firstSum += int64(firstFour[i])
	}

	var lastSum int64
	for i := 0; i < len(lastFour); i++ {
		lastSum += int64(lastFour[i])
	}

	r := fmt.Sprintf("%d %d", firstSum, lastSum)
	fmt.Println(r)
}
