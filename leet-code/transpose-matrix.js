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

/**
 * @param {number[][]} A
 * @return {number[][]}
 */
function transpose(A) {
  let matrix = []

  for (let i = 0; i < A.length; i++) {
    for (let j = 0; j < A[i].length; j++) {
      if (!matrix[j]) {
        matrix[j] = []
      }

      matrix[j][i] = A[i][j] 
    }
  }

  return matrix
}

