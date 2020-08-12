package arraychunk

/*
 * Given an array and chunk size, divide the array into many subarrays
 * where each subarray is of length size

 * chunk([1, 2, 3, 4], 2) --> [[ 1, 2], [3, 4]]
 * chunk([1, 2, 3, 4, 5], 2) --> [[ 1, 2], [3, 4], [5]]
 * chunk([1, 2, 3, 4, 5, 6, 7, 8], 3) --> [[ 1, 2, 3], [4, 5, 6], [7, 8]]
 * chunk([1, 2, 3, 4, 5], 4) --> [[ 1, 2, 3, 4], [5]]
 * chunk([1, 2, 3, 4, 5], 10) --> [[ 1, 2, 3, 4, 5]]
 */

// Chunk ...
func Chunk(list []int, size int) [][]int {
	var result [][]int
	var chunk []int

	for i := 0; i < len(list); i++ {
		chunk = append(chunk, list[i])

		if len(chunk) >= size || i == len(list)-1 {
			result = append(result, chunk)
			chunk = []int{}
		}
	}

	return result
}
