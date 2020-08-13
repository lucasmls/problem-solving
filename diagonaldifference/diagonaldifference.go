package diagonaldifference

/**
 * Given a square matrix, calculate the absolute difference between the sums of its diagonals.
 * For example, the square matrix is shown below:
 * |1 2 3|
 * |4 5 6|
 * |9 8 9|

 * The left-to-right diagonal = 1 + 5 + 9 = 15.
 * The right to left diagonal = 3 + 5 + 9 = 17.
 * Their absolute difference is |15 - 17| = 2.

 * @link https://www.hackerrank.com/challenges/diagonal-difference/problem
 */

// Abs ...
func Abs(x int32) int32 {
	if x < 0 {
		return -x
	}
	return x
}

// DiagonalDifference ...
func DiagonalDifference(arr [][]int32) int32 {
	var leftDiagonal int32
	var rightDiagonal int32

	for i := 0; i < len(arr); i++ {
		row := arr[i]

		actualEl := row[i]
		againstEl := row[len(row)-1-i]

		leftDiagonal += actualEl
		rightDiagonal += againstEl
	}

	return Abs(leftDiagonal - rightDiagonal)
}
