function isToeplitzMatrix(matrix) {
  for (let i = 0; i < matrix.length -1; i++) {
    const list = matrix[i];
    const nextList = matrix[i + 1]

    for (let j = 0; j < list.length -1; j++) {
      const element = list[j]
      const nextElement = nextList[j + 1]

      if (element != nextElement) {
        return false
      }
    }
  }

  return true
}

const matrix = [
  [1,2,3,4],
  [5,1,2,3],
  [9,5,1,2]
]

console.log(isToeplitzMatrix(matrix))