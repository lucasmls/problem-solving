function diagonalDifference(arr) {
  let diagonalX = 0
  let diagonalY = 0

  for (let i = 0; i < arr.length; i++) {
    const lineSize = arr[i].length

    const actualValue = arr[i][i]
    const againstValue = arr[i][lineSize - (i + 1)]

    diagonalX += actualValue
    diagonalY += againstValue
  }

  return Math.abs(diagonalX - diagonalY)
}

const result = diagonalDifference([
  [11, 2, 4],
  [4, 5, 6],
  [10, 8, -12],
])

console.log(result)
