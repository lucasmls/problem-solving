package countofmatchesintournament

/**
 * You are given an integer n, the number of teams in a tournament that has strange rules:
 * - If the current number of teams is even, each team gets paired with another team. A total of n / 2 matches are played, and n / 2 teams advance to the next round.
 * - If the current number of teams is odd, one team randomly advances in the tournament, and the rest gets paired. A total of (n - 1) / 2 matches are played, and (n - 1) / 2 + 1 teams advance to the next round.
 *
 * Input: n = 7
 * Output: 6
 *
 * Input: n = 14
 * Output: 13
 *
 * @link https://leetcode.com/problems/count-of-matches-in-tournament/
 */

func numberOfMatches(n int) int {
	result := 0

	for n > 1 {
		isEven := n%2 == 0

		if isEven {
			result += n / 2
			n = n / 2

			continue
		}

		result += (n - 1) / 2
		n = ((n - 1) / 2) + 1
	}

	return result
}
