package findthehighestaltitude

/**
 * There is a biker going on a road trip. The road trip consists of n + 1 points at different altitudes. The biker starts his trip on point 0 with altitude equal 0.
 * You are given an integer array gain of length n where gain[i] is the net gain in altitude between points i​​​​​​ and i + 1 for all (0 <= i < n). Return the highest altitude of a point.
 *
 * Input: gain = [-5,1,5,0,-7]
 * Output: 1
 * Explanation: The altitudes are [0,-5,-4,1,1,-6]. The highest is 1.
 *
 * Input: gain = [-4,-3,-2,-1,4,3,2]
 * Output: 0
 * Explanation: The altitudes are [0,-4,-7,-9,-10,-6,-3,-1]. The highest is 0.
 *
 * @link https://leetcode.com/problems/find-the-highest-altitude/
 */

func largestAltitude(gain []int) int {
	gainHistory := []int{}
	for i := 0; i < len(gain); i++ {
		lastGain := 0
		if i > 0 {
			lastGain = gainHistory[i-1]
		}

		gainHistory = append(gainHistory, lastGain+gain[i])
	}

	result := 0

	for i := 0; i < len(gainHistory); i++ {
		v := gainHistory[i]
		if v > result {
			result = v
		}
	}

	return result
}
