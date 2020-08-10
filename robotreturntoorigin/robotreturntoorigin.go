package robotreturntoorigin

import (
	"fmt"
)

/**
 * There is a robot starting at position (0, 0), the origin, on a 2D plane. Given a sequence of its moves,
 * judge if this robot ends up at (0, 0) after it completes its moves.
 * The move sequence is represented by a string, and the character moves[i] represents its ith move.
 * Valid moves are R (right), L (left), U (up), and D (down). If the robot returns to the origin after it finishes all of its moves, return true. Otherwise, return false.
 *
 * Input: "UD"
 * Output: true
 * Explanation: The robot moves up once, and then down once. All moves have the same magnitude, so it ended up at the origin where it started. Therefore, we return true.
 *
 * @link https://leetcode.com/problems/robot-return-to-origin/
 */

// CartesianPlane ...
type CartesianPlane struct {
	X int
	Y int
}

// JudgeCircle ...
func JudgeCircle(moves string) bool {
	plane := CartesianPlane{X: 0, Y: 0}

	for i := 0; i < len(moves); i++ {
		move := fmt.Sprintf("%c", moves[i])

		switch move {
		case "R":
			plane.X++
			break
		case "L":
			plane.X--
			break
		case "U":
			plane.Y++
			break
		case "D":
			plane.Y--
			break
		}
	}

	if plane.X == 0 && plane.Y == 0 {
		return true
	}

	return false
}
