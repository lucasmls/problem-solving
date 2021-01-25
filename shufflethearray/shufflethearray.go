package shufflethearray

// https://leetcode.com/problems/shuffle-the-array/submissions/

func shuffle(nums []int, n int) []int {
	result := []int{}

	for i := 0; i < len(nums)/2; i++ {
		x := nums[i]
		y := nums[i+n]

		result = append(result, x)
		result = append(result, y)
	}

	return result
}
