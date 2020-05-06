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

/**
 * @param {string} moves
 * @return {boolean}
 */
function judgeCircle(moves) {
  let [x, y] = [0, 0]

  for (const move of moves) {
    switch (move) {
      case "R":
        x += 1
        break

      case "L":
        x -= 1
        break

      case "U":
        y += 1
        break

      case "D":
        y -= 1
        break
    }
  }

  if (x == 0 && y == 0)
    return true

  return false
}

console.log(judgeCircle("RLUURDDDLU"))
