function simpleArraySum(ar) {
  return ar.reduce((acc, item) => acc += item)
}

const result = simpleArraySum([1, 2, 3, 4, 10, 11])
console.log(result)