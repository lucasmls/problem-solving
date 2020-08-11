package pascalstriangle

/**
 * Given a non-negative integer numRows, generate the first numRows of Pascal's triangle.
 * In Pascal's triangle, each number is the sum of the two numbers directly above it.
 *
 * @link https://leetcode.com/problems/pascals-triangle/
 */

// Generate ...
func Generate(numRows int) [][]int {
	triangle := make([][]int, numRows)

	for i := 1; i <= numRows; i++ {
		row := make([]int, i)
		row[0] = 1
		row[len(row)-1] = 1

		if i > 2 {
			previousRow := triangle[i-2]

			for j := 1; j < len(row)-1; j++ {
				sum := previousRow[j] + previousRow[j-1]
				row[j] = sum
			}
		}

		triangle[i-1] = row
	}

	return triangle
}
