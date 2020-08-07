package transposematrix

/**
 * Given a matrix A, return the transpose of A.
 * The transpose of a matrix is the matrix flipped over it's main diagonal, switching the row and column indices of the matrix.
 *
 * Example:
 * Input: [[1,2,3],[4,5,6],[7,8,9]]
 * Output: [[1,4,7],[2,5,8],[3,6,9]]
 *
 * @link https://leetcode.com/problems/transpose-matrix/
 */

// Transpose the given matrix
func Transpose(A [][]int) [][]int {
	var transposedMatrix [][]int

	for i := 0; i < len(A[0]); i++ {
		transposedMatrix = append(transposedMatrix, make([]int, len(A)))
	}

	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A[i]); j++ {
			transposedMatrix[j][i] = A[i][j]
		}
	}

	return transposedMatrix
}
