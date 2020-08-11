package toeplitzmatrix

/**
 * A matrix is Toeplitz if every diagonal from top-left to bottom-right has the same element.
 * Now given an M x N matrix, return True if and only if the matrix is Toeplitz.
 *
 * @link https://leetcode.com/problems/toeplitz-matrix/
 */

func isToeplitzMatrix(matrix [][]int) bool {
	for i := 0; i < len(matrix)-1; i++ {
		row := matrix[i]
		nextRow := matrix[i+1]

		for j := 0; j < len(row)-1; j++ {
			el := row[j]
			nextEl := nextRow[j+1]

			if el != nextEl {
				return false
			}
		}
	}

	return true
}
