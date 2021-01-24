package shufflestring

// https://leetcode.com/problems/shuffle-string/

func restoreString(s string, indices []int) string {
	shuffledSlice := make([]string, len(s))

	for i := 0; i < len(s); i++ {
		v := string(s[i])
		pos := indices[i]

		shuffledSlice[pos] = v
	}

	result := ""
	for i := 0; i < len(shuffledSlice); i++ {
		result += string(shuffledSlice[i])
	}

	return result
}
