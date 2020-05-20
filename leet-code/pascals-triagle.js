/**
 * Given a non-negative integer numRows, generate the first numRows of Pascal's triangle.
 * In Pascal's triangle, each number is the sum of the two numbers directly above it.
 * 
 * @link https://leetcode.com/problems/pascals-triangle/
 */

/**
 * @param {number} numRows
 * @return {number[][]}
 */
function generate(numRows) {
  let triangle = []

  for (let i = 1; i <= numRows; i++) {
    let row = Array(i).fill(null)
    row[0] = 1
    row[row.length -1] = 1

    if (i > 2) {
      const previousRow = triangle[i -2]
      for (let j = 1; j < row.length -1; j++) {
        const sum = previousRow[j] + previousRow[j -1]
        row[j] = sum
      }
    }

    triangle.push(row)
  }

  return triangle
}

console.log(generate(10))
