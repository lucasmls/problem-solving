package singlenumber

// SingleNumber ...
func SingleNumber(nums []int) int {
	setSum := 0
	numsSum := 0

	numsSet := map[int]bool{}
	for _, num := range nums {
		numsSet[num] = true

		numsSum += num
	}

	for num := range numsSet {
		setSum += num
	}

	return (2 * setSum) - numsSum
}
