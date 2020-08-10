package greatestnumberofcandies

/**
 * Given the array candies and the integer extraCandies, where candies[i] represents the number
 * of candies that the ith kid has.
 * For each kid check if there is a way to distribute extraCandies among the kids such that he or
 * she can have the greatest number of candies among them. Notice that multiple kids can have the
 * greatest number of candies.
 *
 * @link https://leetcode.com/problems/kids-with-the-greatest-number-of-candies/
 */

// KidsWithCandies ...
func KidsWithCandies(candies []int, extraCandies int) []bool {
	maxCandy := 0

	for i := 0; i < len(candies); i++ {
		candy := candies[i]

		if candy >= maxCandy {
			maxCandy = candy
		}
	}

	var result []bool
	for i := 0; i < len(candies); i++ {
		candy := candies[i]
		if (candy + extraCandies) >= maxCandy {
			result = append(result, true)
			continue
		}

		result = append(result, false)
	}

	return result
}
