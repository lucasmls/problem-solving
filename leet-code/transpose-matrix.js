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

