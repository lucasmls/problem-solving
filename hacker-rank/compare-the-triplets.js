function compareTriplets(a, b) {
  let aScore = 0
  let bScore = 0

  a.forEach((item, i) => {
    if (item > b[i]) {
      aScore += 1
      return
    }

    if (item === b[i]) {
      return
    }

    bScore += 1
  })

  return [aScore, bScore]
}


const result = compareTriplets([1, 2, 3], [3, 2, 1])
console.log(result)